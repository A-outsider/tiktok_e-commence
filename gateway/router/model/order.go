package model

import (
	"github.com/cloudwego/hertz/pkg/route"
	"gomall/gateway/controller/order"
	"gomall/gateway/middleware"
)

func RegisterOrder(r *route.RouterGroup) {
	orderApi := order.NewApi()

	r.Use(middleware.Auth())

	r.POST("", orderApi.PlaceOrder) // 创建订单
	r.GET("", orderApi.ListOrder)   // 获取订单列表
	r.GET("/seller", orderApi.ListOrderFromSeller)
	r.PUT("/shipped", orderApi.MarkOrderShipped)
	r.PUT("/completed", orderApi.MarkOrderCompleted)
}
