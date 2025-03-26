package handler

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"gomall/services/seckill/config"
	"gomall/services/seckill/dal/cache"
	"gomall/services/seckill/dal/db"
	"gomall/services/seckill/dal/model"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

// SeckillServiceImpl 秒杀服务实现
type SeckillServiceImpl struct {
	// 限流器
	limiter *rate.Limiter
	// 限购检查锁
	buyLimitLock sync.Mutex
}

// NewSeckillServiceImpl 创建秒杀服务实现
func NewSeckillServiceImpl() *SeckillServiceImpl {
	// 从配置中获取限流参数
	limiterConfig := config.GetConf().Limiter
	// 创建限流器
	limiter := rate.NewLimiter(rate.Limit(limiterConfig.RatePerSecond), limiterConfig.Burst)

	return &SeckillServiceImpl{
		limiter: limiter,
	}
}

// CreateSeckillProduct 创建秒杀商品
func (s *SeckillServiceImpl) CreateSeckillProduct(ctx context.Context, product *model.SeckillProduct) error {
	// 生成秒杀商品ID
	product.SPID = fmt.Sprintf("sp_%d_%s", time.Now().Unix(), product.PID)

	// 保存到数据库
	err := db.CreateSeckillProduct(ctx, product)
	if err != nil {
		zap.L().Error("Failed to create seckill product in database", zap.Error(err))
		return err
	}

	// 初始化Redis库存
	err = cache.InitProductStock(ctx, product.SPID, product.Stock)
	if err != nil {
		zap.L().Error("Failed to initialize product stock in Redis", zap.Error(err))
		return err
	}

	// 缓存商品信息
	err = cache.CacheProductInfo(ctx, product)
	if err != nil {
		zap.L().Error("Failed to cache product info", zap.Error(err))
		// 非致命错误，可以继续
	}

	// 添加到布隆过滤器
	err = cache.AddProductToBloomFilter(ctx, product.SPID)
	if err != nil {
		zap.L().Error("Failed to add product to bloom filter", zap.Error(err))
		// 非致命错误，可以继续
	}

	// 如果是已激活且在活动时间内，添加到活动列表
	now := time.Now()
	if product.IsActive && now.After(product.StartTime) && now.Before(product.EndTime) {
		err = cache.AddToActiveList(ctx, product.SPID)
		if err != nil {
			zap.L().Error("Failed to add product to active list", zap.Error(err))
			// 非致命错误，可以继续
		}
	}

	return nil
}

// GetSeckillProduct 获取秒杀商品信息
func (s *SeckillServiceImpl) GetSeckillProduct(ctx context.Context, spid string) (*model.SeckillProduct, error) {
	// 先尝试从缓存获取
	product, err := cache.GetCachedProductInfo(ctx, spid)
	if err == nil {
		return product, nil
	}

	// 缓存不存在，从数据库获取
	product, err = db.GetSeckillProductBySPID(ctx, spid)
	if err != nil {
		return nil, err
	}

	// 更新缓存
	_ = cache.CacheProductInfo(ctx, product)

	return product, nil
}

// ListActiveSeckillProducts 获取活动中的秒杀商品列表
func (s *SeckillServiceImpl) ListActiveSeckillProducts(ctx context.Context) ([]*model.SeckillProduct, error) {
	// 从Redis获取活动商品ID列表
	spids, err := cache.GetActiveList(ctx)
	if err != nil {
		// Redis获取失败，回退到数据库
		return db.ListActiveSeckillProducts(ctx)
	}

	products := make([]*model.SeckillProduct, 0, len(spids))
	for _, spid := range spids {
		product, err := s.GetSeckillProduct(ctx, spid)
		if err != nil {
			zap.L().Error("Failed to get product", zap.String("spid", spid), zap.Error(err))
			continue
		}
		products = append(products, product)
	}

	return products, nil
}

// DoSeckill 执行秒杀
func (s *SeckillServiceImpl) DoSeckill(ctx context.Context, spid string, uid string, quantity int64) (string, error) {
	// 限流检查
	if !s.limiter.Allow() {
		return "", errors.New("系统繁忙，请稍后再试")
	}

	// 使用布隆过滤器快速检查商品是否存在
	exists, err := cache.ExistsInBloomFilter(ctx, spid)
	if err != nil || !exists {
		return "", errors.New("秒杀商品不存在")
	}

	// 获取商品信息
	product, err := s.GetSeckillProduct(ctx, spid)
	if err != nil {
		return "", err
	}

	// 检查活动状态
	now := time.Now()
	if !product.IsActive {
		return "", errors.New("秒杀活动未激活")
	}
	if now.Before(product.StartTime) {
		return "", errors.New("秒杀活动未开始")
	}
	if now.After(product.EndTime) {
		return "", errors.New("秒杀活动已结束")
	}

	// 检查购买限制
	s.buyLimitLock.Lock()
	canBuy, err := cache.CheckUserBuyLimit(ctx, spid, uid, product.LimitPerUser)
	if err != nil {
		s.buyLimitLock.Unlock()
		return "", err
	}
	if !canBuy {
		s.buyLimitLock.Unlock()
		return "", errors.New("超过购买限制")
	}

	// 预扣减Redis库存
	success, err := cache.DecrProductStock(ctx, spid, quantity)
	if err != nil || !success {
		s.buyLimitLock.Unlock()
		if err != nil {
			return "", err
		}
		return "", errors.New("库存不足")
	}

	// 增加用户购买记录
	err = cache.IncrUserBuyCount(ctx, spid, uid, product.EndTime.Sub(now)+time.Hour)
	s.buyLimitLock.Unlock()
	if err != nil {
		// 购买记录增加失败，回滚库存
		_ = cache.IncrProductStock(ctx, spid, quantity)
		return "", err
	}

	// 生成流水ID
	flowID := fmt.Sprintf("flow_%s_%d", spid, time.Now().UnixNano())

	// 创建并保存库存流水到Redis
	flow := &model.InventoryFlow{
		FlowID:     flowID,
		SPID:       spid,
		UID:        uid,
		OrderID:    "",
		OpType:     1, // 预扣
		Quantity:   quantity,
		Status:     0, // 处理中
		RetryCount: 0,
		LockToken:  uuid.New().String(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err = cache.SaveInventoryFlow(ctx, flow)
	if err != nil {
		// 保存流水失败，回滚库存
		_ = cache.IncrProductStock(ctx, spid, quantity)
		return "", err
	}

	// 异步更新数据库（这里可以使用RocketMQ发送消息）
	// TODO: 发送RocketMQ消息

	return flowID, nil
}

// ConfirmSeckillOrder 确认秒杀订单
func (s *SeckillServiceImpl) ConfirmSeckillOrder(ctx context.Context, flowID string, orderId string) error {
	// 获取流水信息
	flow, err := cache.GetInventoryFlow(ctx, flowID)
	if err != nil {
		return err
	}

	// 确认库存扣减
	err = db.ConfirmInventory(ctx, flowID, orderId)
	if err != nil {
		return err
	}

	// 更新Redis中的流水状态
	flow.Status = 1 // 成功
	flow.OrderID = orderId
	flow.OpType = 2 // 确认
	flow.UpdatedAt = time.Now()

	return cache.SaveInventoryFlow(ctx, flow)
}

// CancelSeckill 取消秒杀
func (s *SeckillServiceImpl) CancelSeckill(ctx context.Context, flowID string) error {
	// 获取流水信息
	flow, err := cache.GetInventoryFlow(ctx, flowID)
	if err != nil {
		return err
	}

	// 回滚库存
	err = db.RollbackInventory(ctx, flowID)
	if err != nil {
		return err
	}

	// 回滚Redis库存
	err = cache.IncrProductStock(ctx, flow.SPID, flow.Quantity)
	if err != nil {
		return err
	}

	// 更新Redis中的流水状态
	flow.Status = 2 // 失败
	flow.OpType = 3 // 回滚
	flow.UpdatedAt = time.Now()

	return cache.SaveInventoryFlow(ctx, flow)
}
