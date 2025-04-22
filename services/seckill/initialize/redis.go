package initialize

import (
	"context"
	"fmt"
	"time"

	"gomall/services/seckill/config"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
)

var _rdb *redis.Client

// InitRedis 初始化Redis连接
func InitRedis() {
	// 从配置中获取Redis配置
	redisConf := config.GetConf().Redis

	// 创建Redis客户端
	_rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port),
		Password: redisConf.Password,
		DB:       redisConf.DB,
		PoolSize: redisConf.PoolSize,
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := _rdb.Ping(ctx).Result()
	if err != nil {
		klog.Fatalf("failed to connect redis: %v", err)
	}

	klog.Infof("Redis connected: %s:%d", redisConf.Host, redisConf.Port)
}

// GetRedis 获取Redis客户端
func GetRedis() *redis.Client {
	return _rdb
}
