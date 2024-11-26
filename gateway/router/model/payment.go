package model

import (
	"github.com/cloudwego/hertz/pkg/route"
	"gomall/gateway/controller/payment"
	"gomall/gateway/middleware"
)

func RegisterPayment(r *route.RouterGroup) {
	paymentApi := payment.NewApi()

	r.GET("/createPay", middleware.Auth(), paymentApi.CreatePayment)
	r.GET("/callback", paymentApi.PayCallback)
	r.POST("/notify", paymentApi.PayNotify)
}
