package utils

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
)

var (
	// 使用一个简单的计数器作为ID的一部分
	counter int64 = 0
	// 服务启动时间戳，作为ID的一部分
	startTimeNano = time.Now().UnixNano()
)

// GenerateUniqueID 生成唯一ID
func GenerateUniqueID() string {
	// 增加计数器
	count := atomic.AddInt64(&counter, 1)

	// 获取当前时间戳
	timestamp := time.Now().UnixNano()

	// 组合时间戳、启动时间和计数器
	return fmt.Sprintf("%d%d%d", timestamp, startTimeNano%10000, count)
}

// GenerateOrderID 生成订单ID
func GenerateOrderID(prefix string) string {
	// 增加计数器
	count := atomic.AddInt64(&counter, 1)

	// 获取当前时间戳
	timestamp := time.Now().UnixMilli()

	// 组合前缀、时间戳和计数器
	return fmt.Sprintf("%s%d%d", prefix, timestamp, count)
}

// GenerateIdempotentToken 生成幂等性Token
func GenerateIdempotentToken() string {
	return uuid.NewString()
}
