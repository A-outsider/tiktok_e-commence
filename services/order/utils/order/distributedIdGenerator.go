package order

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// DistributedIdGenerator 是一个分布式 ID 生成器
type DistributedIdGenerator struct {
	EPOCH         int64 // 开始时间的时间戳 (毫秒)
	NODE_BITS     int64 // 节点 ID 所占的位数
	SEQUENCE_BITS int64 // 序列号所占的位数
	nodeID        int64 // 节点 ID
	lastTimestamp int64 // 上一次生成 ID 的时间戳
	sequence      int64 // 序列号
	mutex         sync.Mutex
}

// NewDistributedIdGenerator 初始化分布式 ID 生成器
func NewDistributedIdGenerator(nodeID int64) *DistributedIdGenerator {
	return &DistributedIdGenerator{
		EPOCH:         1609459200000, // 2021-01-01 00:00:00 UTC
		NODE_BITS:     5,
		SEQUENCE_BITS: 7,
		nodeID:        nodeID,
		lastTimestamp: -1,
		sequence:      0,
	}
}

// GenerateId 生成唯一的分布式 ID
func (g *DistributedIdGenerator) GenerateId() (int64, error) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	timestamp := time.Now().UnixMilli() - g.EPOCH
	if timestamp < g.lastTimestamp {
		return 0, fmt.Errorf("clock moved backwards. refusing to generate ID")
	}

	if timestamp == g.lastTimestamp {
		g.sequence = (g.sequence + 1) & ((1 << g.SEQUENCE_BITS) - 1)
		if g.sequence == 0 {
			timestamp = g.tilNextMillis(g.lastTimestamp)
		}
	} else {
		g.sequence = 0
	}

	g.lastTimestamp = timestamp
	return (timestamp << (g.NODE_BITS + g.SEQUENCE_BITS)) | (g.nodeID << g.SEQUENCE_BITS) | g.sequence, nil
}

// tilNextMillis 等待直到下一毫秒
func (g *DistributedIdGenerator) tilNextMillis(lastTimestamp int64) int64 {
	timestamp := time.Now().UnixMilli() - g.EPOCH
	for timestamp <= lastTimestamp {
		timestamp = time.Now().UnixMilli() - g.EPOCH
	}
	return timestamp
}

// OrderIdGeneratorManager 订单 ID 生成管理器，使用单例模式
type OrderIdGeneratorManager struct {
	idGenerator *DistributedIdGenerator
}

// 全局的 OrderIdGeneratorManager 实例
var (
	instance *OrderIdGeneratorManager
	once     sync.Once
)

// GetOrderIdGeneratorManager 获取全局唯一的订单 ID 生成管理器实例
func GetOrderIdGeneratorManager() *OrderIdGeneratorManager {
	once.Do(func() {
		// 使用随机 nodeID 创建 DistributedIdGenerator
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		nodeID := r.Int63n(1 << 5) // 生成一个 0 到 31 之间的随机 nodeID
		instance = &OrderIdGeneratorManager{
			idGenerator: NewDistributedIdGenerator(nodeID),
		}
	})
	return instance
}

// GenerateId 生成订单 ID，传入 userId 后生成全局唯一的订单 ID
func (o *OrderIdGeneratorManager) GenerateId(userId string) (string, error) {
	id, err := o.idGenerator.GenerateId()
	if err != nil {
		return "", fmt.Errorf("error generating distributedIdGeneratorID")
	}
	parseInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d%d", id, parseInt%1000000), nil
}
