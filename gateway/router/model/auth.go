package model

import (
	"github.com/cloudwego/hertz/pkg/route"
	"gomall/gateway/controller/auth"
)

func RegisterAuth(r *route.RouterGroup) {
	authApi := auth.NewApi()

	r.POST("/register", authApi.Register)
	r.POST("/login/phone_code", authApi.LoginByCode)
	r.POST("/login/password", authApi.LoginByPwd)

	r.POST("/phone", authApi.SendPhoneCode)
	r.POST("/email", authApi.SendEmailCode)

	r.POST("/refresh_token", authApi.RefreshToken)
	r.POST("/photo_captcha", authApi.ShowPhotoCaptcha)

	// 管理员接口 TODO : 暂时写这
	r.POST("/user_role", authApi.ModifyUserToSeller)
}
