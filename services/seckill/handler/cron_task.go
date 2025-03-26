package handler

import (
	"context"
	"time"

	"gomall/services/seckill/config"
	"gomall/services/seckill/dal/cache"
	"gomall/services/seckill/dal/db"

	"go.uber.org/zap"
)

// CronTaskManager 定时任务管理器
type CronTaskManager struct {
	seckillService *SeckillServiceImpl
	quit           chan struct{}
}

// NewCronTaskManager 创建定时任务管理器
func NewCronTaskManager(seckillService *SeckillServiceImpl) *CronTaskManager {
	return &CronTaskManager{
		seckillService: seckillService,
		quit:           make(chan struct{}),
	}
}

// Start 启动定时任务
func (c *CronTaskManager) Start() {
	go c.runPendingFlowsTask()
	go c.runExpiredOrdersTask()
}

// Stop 停止定时任务
func (c *CronTaskManager) Stop() {
	close(c.quit)
}

// runPendingFlowsTask 处理待处理的库存流水任务
func (c *CronTaskManager) runPendingFlowsTask() {
	interval := time.Duration(config.GetConf().Seckill.CheckInterval) * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.processPendingFlows()
		case <-c.quit:
			zap.L().Info("Pending flows task stopped")
			return
		}
	}
}

// runExpiredOrdersTask 处理过期订单任务
func (c *CronTaskManager) runExpiredOrdersTask() {
	interval := time.Duration(config.GetConf().Seckill.CheckInterval) * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.processExpiredOrders()
		case <-c.quit:
			zap.L().Info("Expired orders task stopped")
			return
		}
	}
}

// processPendingFlows 处理待处理的库存流水
func (c *CronTaskManager) processPendingFlows() {
	ctx := context.Background()

	// 获取待处理的流水记录
	flows, err := db.ListPendingFlows(ctx, 100) // 每次处理100条
	if err != nil {
		zap.L().Error("Failed to list pending flows", zap.Error(err))
		return
	}

	zap.L().Info("Processing pending flows", zap.Int("count", len(flows)))

	for _, flow := range flows {
		// 如果流水创建时间超过订单TTL时间，则回滚库存
		if time.Since(flow.CreatedAt) > time.Duration(config.GetConf().Seckill.OrderTTL)*time.Second {
			zap.L().Info("Rolling back expired flow", zap.String("flow_id", flow.FlowID))
			err := c.seckillService.CancelSeckill(ctx, flow.FlowID)
			if err != nil {
				zap.L().Error("Failed to rollback inventory", zap.String("flow_id", flow.FlowID), zap.Error(err))
			}
		}
	}
}

// processExpiredOrders 处理过期订单
func (c *CronTaskManager) processExpiredOrders() {
	// 这个方法主要是与订单服务配合，检查已生成订单但未支付的情况
	// 由于秒杀模块已经有了库存流水的处理，这里主要是双重保障机制

	// 实际项目中，可能需要调用订单服务获取未支付的秒杀订单
	// 然后遍历这些订单，检查是否过期，过期则取消订单并回滚库存

	zap.L().Info("Processing expired orders - placeholder for integration with order service")
}

// verifyInventoryConsistency 验证Redis与MySQL库存一致性
func (c *CronTaskManager) verifyInventoryConsistency() {
	ctx := context.Background()

	// 获取活动中的商品
	products, err := db.ListActiveSeckillProducts(ctx)
	if err != nil {
		zap.L().Error("Failed to list active products", zap.Error(err))
		return
	}

	for _, product := range products {
		// 获取Redis库存
		redisStock, err := cache.GetProductStock(ctx, product.SPID)
		if err != nil {
			zap.L().Error("Failed to get Redis stock", zap.String("spid", product.SPID), zap.Error(err))
			continue
		}

		// 对比Redis库存与MySQL库存
		if redisStock != product.Stock {
			zap.L().Warn("Inventory inconsistency detected",
				zap.String("spid", product.SPID),
				zap.Int64("mysql_stock", product.Stock),
				zap.Int64("redis_stock", redisStock))

			// 根据业务策略决定如何处理一致性问题
			// 例如：以MySQL库存为准，更新Redis库存
			err = cache.InitProductStock(ctx, product.SPID, product.Stock)
			if err != nil {
				zap.L().Error("Failed to update Redis stock", zap.String("spid", product.SPID), zap.Error(err))
			}
		}
	}
}
