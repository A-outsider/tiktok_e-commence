package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gomall/services/seckill/initialize"

	"github.com/redis/go-redis/v9"
)

// Redis缓存前缀
const (
	ProductStockPrefix   = "seckill:stock:"
	UserBoughtPrefix     = "seckill:user:"
	BloomFilterKey       = "seckill:bloom"
	ProductInfoPrefix    = "seckill:product:"
	ActiveProductListKey = "seckill:active"
	LockPrefix           = "seckill:lock:"
	InventoryFlowPrefix  = "seckill:flow:"
	IdempotentPrefix     = "seckill:idemp:"
)

// 错误定义
var (
	ErrKeyNotExist  = errors.New("key not exist")
	ErrLockFailed   = errors.New("failed to acquire lock")
	ErrUnlockFailed = errors.New("failed to release lock")
)

// IdempotentInfo 幂等信息
type IdempotentInfo struct {
	SPID      int64     `json:"spid"`       // 秒杀商品ID
	UID       int64     `json:"uid"`        // 用户ID
	Quantity  int32     `json:"quantity"`   // 购买数量
	FlowID    string    `json:"flow_id"`    // 流水ID
	Status    int       `json:"status"`     // 状态：0-处理中 1-成功 2-失败
	OrderID   string    `json:"order_id"`   // 订单ID
	CreatedAt time.Time `json:"created_at"` // 创建时间
	ExpiredAt time.Time `json:"expired_at"` // 过期时间
}

// GetRedis 获取Redis客户端
func GetRedis() *redis.Client {
	return initialize.GetRedis()
}

// SetProductStock 设置商品库存
func SetProductStock(ctx context.Context, spid int64, stock int32) error {
	key := fmt.Sprintf("%s%d", ProductStockPrefix, spid)
	return GetRedis().Set(ctx, key, stock, 0).Err()
}

// GetProductStock 获取商品库存
func GetProductStock(ctx context.Context, spid int64) (int32, error) {
	key := fmt.Sprintf("%s%d", ProductStockPrefix, spid)
	val, err := GetRedis().Get(ctx, key).Int()
	if err != nil {
		if err == redis.Nil {
			return 0, ErrKeyNotExist
		}
		return 0, err
	}
	return int32(val), nil
}

// DecrProductStock 减少商品库存
func DecrProductStock(ctx context.Context, spid int64, quantity int32) (bool, error) {
	key := fmt.Sprintf("%s%d", ProductStockPrefix, spid)

	// 使用Lua脚本原子性地检查库存并扣减
	script := `
	local current = tonumber(redis.call('get', KEYS[1]))
	if current == nil then return -1 end
	if current < tonumber(ARGV[1]) then return 0 end
	redis.call('decrby', KEYS[1], ARGV[1])
	return 1
	`

	res, err := GetRedis().Eval(ctx, script, []string{key}, quantity).Int()
	if err != nil {
		return false, err
	}

	if res == -1 {
		return false, ErrKeyNotExist
	}

	return res == 1, nil
}

// IncrProductStock 增加商品库存
func IncrProductStock(ctx context.Context, spid int64, quantity int32) error {
	key := fmt.Sprintf("%s%d", ProductStockPrefix, spid)
	return GetRedis().IncrBy(ctx, key, int64(quantity)).Err()
}

// HasUserBought 检查用户是否已购买
func HasUserBought(ctx context.Context, spid int64, uid int64) (bool, error) {
	key := fmt.Sprintf("%s%d:%d", UserBoughtPrefix, spid, uid)
	exists, err := GetRedis().Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

// MarkUserBought 标记用户已购买
func MarkUserBought(ctx context.Context, spid int64, uid int64) error {
	key := fmt.Sprintf("%s%d:%d", UserBoughtPrefix, spid, uid)
	return GetRedis().Set(ctx, key, 1, 24*time.Hour).Err()
}

// AddToActiveList 添加商品到活动列表
func AddToActiveList(ctx context.Context, spid string) error {
	return GetRedis().SAdd(ctx, ActiveProductListKey, spid).Err()
}

// SaveIdempotentInfo 保存幂等信息
func SaveIdempotentInfo(ctx context.Context, idempToken string, info *IdempotentInfo) error {
	key := fmt.Sprintf("%s%s", IdempotentPrefix, idempToken)
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}

	// 计算过期时间
	var expiration time.Duration
	if !info.ExpiredAt.IsZero() {
		expiration = time.Until(info.ExpiredAt)
		if expiration < 0 {
			expiration = 24 * time.Hour // 默认24小时
		}
	} else {
		expiration = 24 * time.Hour // 默认24小时
	}

	return GetRedis().Set(ctx, key, data, expiration).Err()
}

// GetIdempotentInfo 获取幂等信息
func GetIdempotentInfo(ctx context.Context, idempToken string) (*IdempotentInfo, error) {
	key := fmt.Sprintf("%s%s", IdempotentPrefix, idempToken)
	data, err := GetRedis().Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrKeyNotExist
		}
		return nil, err
	}

	var info IdempotentInfo
	if err := json.Unmarshal(data, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

// UpdateIdempotentStatus 更新幂等信息状态
func UpdateIdempotentStatus(ctx context.Context, idempToken string, status int, orderID string) error {
	info, err := GetIdempotentInfo(ctx, idempToken)
	if err != nil {
		return err
	}

	info.Status = status
	if orderID != "" {
		info.OrderID = orderID
	}

	return SaveIdempotentInfo(ctx, idempToken, info)
}
