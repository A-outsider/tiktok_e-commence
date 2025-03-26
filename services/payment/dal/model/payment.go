package model

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	PayId       string    `json:"pay_id" gorm:"column:pay_id; index" ` // 支付ID
	OrderID     string    `gorm:"index;not null"`                      // 订单ID
	Amount      float64   `gorm:"type:decimal(10,2)"`                  // 实际支付金额
	Currency    string    `gorm:"size:3"`                              // 货币类型
	Status      int       // 0-待支付 1-成功 2-失败 3-退款
	PaymentTime time.Time // 支付时间
	Channel     string    // 支付渠道
	Identifier  string    `json:"identifier" gorm:"unique"` // 幂等号标识

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
