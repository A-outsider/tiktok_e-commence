package model

import (
	"time"

	"gorm.io/gorm"
)

// InventoryFlow 库存流水表
type InventoryFlow struct {
	FlowID     string `gorm:"primaryKey"`     // 流水ID：flow_秒杀商品ID_操作时间戳
	SPID       string `gorm:"index;not null"` // 关联秒杀商品
	UID        string `gorm:"index;not null"` // 用户ID
	OrderID    string `gorm:"index"`          // 关联订单（可为空）
	OpType     int    `gorm:"not null"`       // 操作类型：1-预扣 2-确认 3-回滚
	Quantity   int64  `gorm:"not null"`       // 操作数量
	Status     int    `gorm:"default:0"`      // 0-处理中 1-成功 2-失败
	RetryCount int    `gorm:"default:0"`      // 重试次数
	LockToken  string `gorm:"uniqueIndex"`    // 分布式锁令牌

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName 设置表名
func (i *InventoryFlow) TableName() string {
	return "inventory_flow"
}
