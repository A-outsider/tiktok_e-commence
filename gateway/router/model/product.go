package model

import (
	"github.com/cloudwego/hertz/pkg/route"
	"gomall/gateway/controller/product"
	"gomall/gateway/middleware"
)

func RegisterProduct(r *route.RouterGroup) {
	productApi := product.NewApi()

	r.Use(middleware.Auth())

	r.POST("", productApi.AddProduct)
	r.GET("/category", productApi.ListProducts)
	r.GET("/query", productApi.SearchProducts)
	r.DELETE("", productApi.DeleteProduct)
	r.GET("", productApi.GetProduct)
	r.GET("/rankings", productApi.GetRankings)
}
