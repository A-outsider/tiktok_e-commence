package model

import (
	"github.com/cloudwego/hertz/pkg/route"
	"gomall/gateway/controller/auth"
)

func RegisterAuth(r *route.RouterGroup) {
	authApi := auth.NewApi()

	r.POST("/login", authApi.Login)

}
