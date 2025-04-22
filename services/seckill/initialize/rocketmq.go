package initialize

import (
	"context"
	"log"
	"sync"
)

var (
	rocketmqProducer interface{} // 实际项目应改为真实的RocketMQ Producer实例类型
	rocketmqConsumer interface{} // 实际项目应改为真实的RocketMQ Consumer实例类型
	rocketmqLock     sync.Mutex
)

// InitRocketMQ 初始化RocketMQ客户端
func InitRocketMQ(ctx context.Context) error {
	rocketmqLock.Lock()
	defer rocketmqLock.Unlock()

	// 这里应该是真实的RocketMQ客户端初始化代码
	// 由于没有实际的RocketMQ客户端，这里只做模拟
	log.Println("RocketMQ client initialized")
	return nil
}

// GetRocketMQProducer 获取RocketMQ生产者
func GetRocketMQProducer() interface{} {
	return rocketmqProducer
}

// GetRocketMQConsumer 获取RocketMQ消费者
func GetRocketMQConsumer() interface{} {
	return rocketmqConsumer
}

// CloseRocketMQ 关闭RocketMQ连接
func CloseRocketMQ() {
	rocketmqLock.Lock()
	defer rocketmqLock.Unlock()

	// 这里应该是真实的RocketMQ客户端关闭代码
	// 由于没有实际的RocketMQ客户端，这里只做模拟
	log.Println("RocketMQ client closed")
}
