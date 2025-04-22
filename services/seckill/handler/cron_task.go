package handler

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

// CronTaskManager 定时任务管理器
type CronTaskManager struct {
	service *SeckillServiceImpl
	quit    chan bool
}

// NewCronTaskManager 创建一个新的定时任务管理器
func NewCronTaskManager(service *SeckillServiceImpl) *CronTaskManager {
	return &CronTaskManager{
		service: service,
		quit:    make(chan bool),
	}
}

// Start 启动定时任务
func (c *CronTaskManager) Start() {
	klog.Info("Starting cron tasks")

	// 启动过期流水检查任务（每分钟检查一次）
	go c.runExpiredFlowsTask()

	// 启动库存一致性检查任务（每5分钟检查一次）
	go c.runInventoryConsistencyTask()
}

// Stop 停止所有定时任务
func (c *CronTaskManager) Stop() {
	klog.Info("Stopping cron tasks")
	close(c.quit)
}

// runExpiredFlowsTask 处理过期流水的定时任务
func (c *CronTaskManager) runExpiredFlowsTask() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			klog.Info("Running expired flows check task")

			// 创建一个带超时的上下文
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

			// 调用服务方法处理过期流水
			c.service.CheckExpiredFlows(ctx)

			cancel()

		case <-c.quit:
			klog.Info("Expired flows check task stopped")
			return
		}
	}
}

// runInventoryConsistencyTask 检查库存一致性的定时任务
func (c *CronTaskManager) runInventoryConsistencyTask() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			klog.Info("Running inventory consistency check task")

			// 创建一个带超时的上下文
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

			// 调用服务方法检查库存一致性
			c.service.CheckInventoryConsistency(ctx)

			cancel()

		case <-c.quit:
			klog.Info("Inventory consistency check task stopped")
			return
		}
	}
}
