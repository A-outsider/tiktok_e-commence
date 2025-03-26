package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gomall/services/seckill/dal/model"
	"gomall/services/seckill/initialize"
)

const (
	// 秒杀商品库存缓存前缀
	SeckillStockPrefix = "seckill:stock:"
	// 秒杀商品信息缓存前缀
	SeckillProductPrefix = "seckill:product:"
	// 秒杀活动列表缓存键
	SeckillActiveKey = "seckill:active"
	// 用户购买记录前缀
	SeckillUserBuyPrefix = "seckill:user:buy:"
	// 布隆过滤器键名
	SeckillBloomFilter = "seckill:bloom:filter"
	// 库存流水前缀
	InventoryFlowPrefix = "inventory:flow:"
	// 库存操作锁前缀
	InventoryLockPrefix = "inventory:lock:"
)

// InitProductStock 初始化秒杀商品库存到Redis
func InitProductStock(ctx context.Context, spid string, stock int64) error {
	key := fmt.Sprintf("%s%s", SeckillStockPrefix, spid)
	return initialize.GetRedis().Set(ctx, key, stock, 0).Err()
}

// GetProductStock 获取秒杀商品库存
func GetProductStock(ctx context.Context, spid string) (int64, error) {
	key := fmt.Sprintf("%s%s", SeckillStockPrefix, spid)
	val, err := initialize.GetRedis().Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(val, 10, 64)
}

// DecrProductStock 使用Lua脚本原子性地扣减库存
func DecrProductStock(ctx context.Context, spid string, quantity int64) (bool, error) {
	// 使用Lua脚本保证原子性
	script := `
		local stockKey = KEYS[1]
		local decrBy = tonumber(ARGV[1])
		
		-- 获取当前库存
		local currentStock = tonumber(redis.call('get', stockKey))
		if currentStock == nil then
			return 0  -- 商品不存在
		end
		
		-- 检查库存是否足够
		if currentStock < decrBy then
			return -1  -- 库存不足
		end
		
		-- 扣减库存
		redis.call('decrby', stockKey, decrBy)
		return 1  -- 扣减成功
	`

	key := fmt.Sprintf("%s%s", SeckillStockPrefix, spid)
	result, err := initialize.GetRedis().Eval(ctx, script, []string{key}, quantity).Int64()
	if err != nil {
		return false, err
	}

	switch result {
	case 1:
		return true, nil // 扣减成功
	case 0:
		return false, errors.New("商品不存在")
	case -1:
		return false, errors.New("库存不足")
	default:
		return false, errors.New("未知错误")
	}
}

// IncrProductStock 增加库存（回滚使用）
func IncrProductStock(ctx context.Context, spid string, quantity int64) error {
	key := fmt.Sprintf("%s%s", SeckillStockPrefix, spid)
	return initialize.GetRedis().IncrBy(ctx, key, quantity).Err()
}

// CheckUserBuyLimit 检查用户是否超过购买限制
func CheckUserBuyLimit(ctx context.Context, spid string, uid string, limit int) (bool, error) {
	key := fmt.Sprintf("%s%s:%s", SeckillUserBuyPrefix, spid, uid)
	count, err := initialize.GetRedis().Get(ctx, key).Int()
	if err != nil && err.Error() != "redis: nil" {
		return false, err
	}

	return count < limit, nil
}

// IncrUserBuyCount 增加用户购买次数
func IncrUserBuyCount(ctx context.Context, spid string, uid string, expiration time.Duration) error {
	key := fmt.Sprintf("%s%s:%s", SeckillUserBuyPrefix, spid, uid)
	pipe := initialize.GetRedis().Pipeline()
	pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, expiration)
	_, err := pipe.Exec(ctx)
	return err
}

// AddProductToBloomFilter 将商品ID添加到布隆过滤器（用于快速判断商品是否参与秒杀）
func AddProductToBloomFilter(ctx context.Context, spid string) error {
	// 这里简化处理，实际应使用专门的布隆过滤器库或服务
	return initialize.GetRedis().SAdd(ctx, SeckillBloomFilter, spid).Err()
}

// ExistsInBloomFilter 检查商品是否在布隆过滤器中
func ExistsInBloomFilter(ctx context.Context, spid string) (bool, error) {
	return initialize.GetRedis().SIsMember(ctx, SeckillBloomFilter, spid).Result()
}

// CacheProductInfo 缓存秒杀商品信息
func CacheProductInfo(ctx context.Context, product *model.SeckillProduct) error {
	key := fmt.Sprintf("%s%s", SeckillProductPrefix, product.SPID)
	data, err := json.Marshal(product)
	if err != nil {
		return err
	}

	// 设置过期时间为活动结束后一小时
	expiration := product.EndTime.Sub(time.Now()) + time.Hour
	return initialize.GetRedis().Set(ctx, key, data, expiration).Err()
}

// GetCachedProductInfo 获取缓存的秒杀商品信息
func GetCachedProductInfo(ctx context.Context, spid string) (*model.SeckillProduct, error) {
	key := fmt.Sprintf("%s%s", SeckillProductPrefix, spid)
	data, err := initialize.GetRedis().Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var product model.SeckillProduct
	if err := json.Unmarshal(data, &product); err != nil {
		return nil, err
	}

	return &product, nil
}

// AddToActiveList 添加商品到活动列表
func AddToActiveList(ctx context.Context, spid string) error {
	return initialize.GetRedis().SAdd(ctx, SeckillActiveKey, spid).Err()
}

// RemoveFromActiveList 从活动列表移除商品
func RemoveFromActiveList(ctx context.Context, spid string) error {
	return initialize.GetRedis().SRem(ctx, SeckillActiveKey, spid).Err()
}

// GetActiveList 获取活动商品列表
func GetActiveList(ctx context.Context) ([]string, error) {
	return initialize.GetRedis().SMembers(ctx, SeckillActiveKey).Result()
}

// AcquireLock 获取分布式锁
func AcquireLock(ctx context.Context, lockKey string, value string, expiration time.Duration) (bool, error) {
	key := fmt.Sprintf("%s%s", InventoryLockPrefix, lockKey)
	return initialize.GetRedis().SetNX(ctx, key, value, expiration).Result()
}

// ReleaseLock 释放分布式锁
func ReleaseLock(ctx context.Context, lockKey string, value string) (bool, error) {
	// 使用Lua脚本确保原子性释放锁
	script := `
		if redis.call('get', KEYS[1]) == ARGV[1] then
			return redis.call('del', KEYS[1])
		else
			return 0
		end
	`

	key := fmt.Sprintf("%s%s", InventoryLockPrefix, lockKey)
	result, err := initialize.GetRedis().Eval(ctx, script, []string{key}, value).Int64()
	if err != nil {
		return false, err
	}

	return result == 1, nil
}

// SaveInventoryFlow 保存库存流水到Redis
func SaveInventoryFlow(ctx context.Context, flow *model.InventoryFlow) error {
	key := fmt.Sprintf("%s%s", InventoryFlowPrefix, flow.FlowID)
	data, err := json.Marshal(flow)
	if err != nil {
		return err
	}

	// 设置适当的过期时间
	return initialize.GetRedis().Set(ctx, key, data, 24*time.Hour).Err()
}

// GetInventoryFlow 获取库存流水
func GetInventoryFlow(ctx context.Context, flowID string) (*model.InventoryFlow, error) {
	key := fmt.Sprintf("%s%s", InventoryFlowPrefix, flowID)
	data, err := initialize.GetRedis().Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var flow model.InventoryFlow
	if err := json.Unmarshal(data, &flow); err != nil {
		return nil, err
	}

	return &flow, nil
}
