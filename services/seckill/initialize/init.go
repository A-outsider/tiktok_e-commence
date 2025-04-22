package initialize

import (
	"context"
	"log"

	"gomall/services/seckill/dal/model"

	"github.com/cloudwego/kitex/pkg/klog"
)

// Init 初始化所有组件
func Init() {
	// 初始化MySQL
	InitMySQL()

	// 自动迁移表结构
	db := GetMysql()
	err := db.AutoMigrate(&model.SeckillProduct{}, &model.InventoryFlow{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	klog.Info("MySQL tables migrated")

	// 初始化Redis
	InitRedis()
	klog.Info("Redis initialized")

	// 初始化RocketMQ
	err = InitRocketMQ(context.Background())
	if err != nil {
		log.Fatalf("failed to initialize RocketMQ: %v", err)
	}
	klog.Info("RocketMQ initialized")
}
