package model

import (
	"gorm.io/gorm"
	"time"
)

// SeckillProduct 秒杀商品表
type SeckillProduct struct {
	SPID         string    `gorm:"primaryKey;comment:秒杀商品ID"` // 格式：sp_时间戳_商品ID
	PID          string    `gorm:"index;not null"`            // 原商品ID
	SeckillPrice float64   `gorm:"type:decimal(10,2)"`        // 秒杀价
	Stock        int64     `gorm:"not null"`                  // 秒杀库存（独立库存）
	StartTime    time.Time `gorm:"index"`                     // 秒杀开始时间
	EndTime      time.Time `gorm:"index"`                     // 秒杀结束时间
	LimitPerUser int       `gorm:"default:1"`                 // 每人限购数量
	IsActive     bool      `gorm:"default:false"`             // 是否激活秒杀
	Version      int64     `gorm:"default:0"`                 // 乐观锁版本号

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName 设置表名
func (s *SeckillProduct) TableName() string {
	return "seckill_product"
}
