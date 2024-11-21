package model

import (
	"github.com/cloudwego/hertz/pkg/route"
	"gomall/gateway/controller/cart"

	"gomall/gateway/middleware"
)

func RegisterCart(r *route.RouterGroup) {
	cartApi := cart.NewApi()

	r.Use(middleware.Auth())

	r.GET("", cartApi.GetCart)
	r.POST("", cartApi.AddItem)
	r.DELETE("", cartApi.EmptyCart)
}
