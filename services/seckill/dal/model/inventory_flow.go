package model

import (
	"time"

	"gorm.io/gorm"
)

// InventoryFlow 库存流水表
type InventoryFlow struct {
	ID         int64  `gorm:"primaryKey;autoIncrement;comment:流水ID"`
	FlowID     string `gorm:"uniqueIndex;comment:流水号;type:varchar(64)"`
	SPID       int64  `gorm:"index;not null;comment:关联秒杀商品ID"`
	UID        int64  `gorm:"index;not null;comment:用户ID"`
	OrderID    string `gorm:"index;comment:关联订单ID;type:varchar(64)"`
	OpType     int    `gorm:"not null;comment:操作类型：1-预扣 2-确认 3-回滚"`
	Quantity   int32  `gorm:"not null;comment:操作数量"`
	Status     int    `gorm:"default:0;comment:状态：0-处理中 1-成功 2-失败"`
	RetryCount int    `gorm:"default:0;comment:重试次数"`
	IdempToken string `gorm:"index;comment:幂等标识;type:varchar(64)"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 设置表名
func (i *InventoryFlow) TableName() string {
	return "inventory_flow"
}
