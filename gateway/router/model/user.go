package model

import (
	"github.com/cloudwego/hertz/pkg/route"
	"gomall/gateway/controller/user"
	"gomall/gateway/middleware"
)

func RegisterUser(r *route.RouterGroup) {
	userApi := user.NewApi()

	r.Use(middleware.Auth())

	r.GET("/info", userApi.GetUserInfo)
	r.PUT("/info", userApi.ModifyUserInfo)
	r.DELETE("", userApi.DeleteUser)
	r.POST("/avatar", userApi.UploadAvatar)

	addr := r.Group("/address")
	addr.GET("", userApi.GetAddressList)
	addr.POST("", userApi.AddAddress)
	addr.PUT("", userApi.ModifyAddress)
	addr.DELETE("", userApi.DeleteAddress)
	addr.POST("/default", userApi.SetDefaultAddress)
}
