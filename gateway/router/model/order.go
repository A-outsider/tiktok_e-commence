package model

import (
	"github.com/cloudwego/hertz/pkg/route"
	"gomall/gateway/controller/order"
	"gomall/gateway/middleware"
)

func RegisterOrder(r *route.RouterGroup) {
	orderApi := order.NewApi()

	r.Use(middleware.Auth())

	r.POST("", orderApi.PlaceOrder)
	r.GET("", orderApi.ListOrder)
}
