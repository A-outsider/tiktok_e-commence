package model

import (
	"time"

	"gorm.io/gorm"
)

// SeckillProduct 秒杀商品表
type SeckillProduct struct {
	ID           int64     `gorm:"primaryKey;autoIncrement;comment:秒杀商品ID"`
	ProductID    int64     `gorm:"index;not null;comment:原商品ID"`
	SeckillPrice float64   `gorm:"type:decimal(10,2);comment:秒杀价"`
	Stock        int32     `gorm:"not null;comment:秒杀库存（独立库存）"`
	StartTime    time.Time `gorm:"index;comment:秒杀开始时间"`
	EndTime      time.Time `gorm:"index;comment:秒杀结束时间"`
	Version      int64     `gorm:"default:0;comment:乐观锁版本号"`
	LimitPerUser int32     `gorm:"default:1;comment:每人限购数量"`
	IsActive     bool      `gorm:"default:false;comment:是否激活"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 设置表名
func (s *SeckillProduct) TableName() string {
	return "seckill_product"
}
