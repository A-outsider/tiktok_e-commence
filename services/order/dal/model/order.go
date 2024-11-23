package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	Oid          string `gorm:"primaryKey; NOT NULL; comment:订单ID" json:"oid"`
	Uid          string `gorm:"NOT NULL; index" json:"uid"`
	UserCurrency string `json:"userCurrency"` // 用户使用货币

	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`

	Pid      string  `gorm:"NOT NULL; index" json:"pid"` // 商品id
	Quantity int64   `json:"quantity"`                   // 商品数量
	Cost     float32 `json:"cost"`                       // 商品花费

	Status int `json:"status"` // 订单状态

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

const (
	// 订单状态常量定义
	OrderStatusPending   = 0 // 待处理（订单已创建，但未支付）
	OrderStatusPaid      = 1 // 已支付（用户完成支付）
	OrderStatusShipped   = 2 // 已发货（订单已发货）
	OrderStatusCompleted = 3 // 已完成（订单完成并确认收货）
	OrderStatusCancelled = 4 // 已取消（订单被取消）
	OrderStatusRefunded  = 5 // 已退款（订单退款成功）
)
