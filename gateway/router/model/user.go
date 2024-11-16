package model

import (
	"github.com/cloudwego/hertz/pkg/route"
	"gomall/gateway/controller/user"
	"gomall/gateway/middleware"
)

func RegisterUser(r *route.RouterGroup) {
	userApi := user.NewApi()

	u := r.Group("/user", middleware.Auth())

	u.GET("/info", userApi.GetUserInfo)
	u.PUT("/info", userApi.ModifyUserInfo)
	u.DELETE("", userApi.DeleteUser)
	u.POST("/avatar", userApi.UploadAvatar)

	addr := u.Group("/address")
	addr.GET("", userApi.GetAddressList)
	addr.POST("", userApi.AddAddress)
	addr.PUT("", userApi.ModifyAddress)
	addr.DELETE("/", userApi.DeleteAddress)
	addr.POST("/default", userApi.SetDefaultAddress)
}
