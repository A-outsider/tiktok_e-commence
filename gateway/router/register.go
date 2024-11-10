package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"gomall/gateway/middleware"
	"gomall/gateway/router/model"
)

func Register(r *server.Hertz) {
	r.Use(middleware.CORS())
	v1 := r.Group("/api/v1")

	// 分模块
	model.RegisterUser(v1)
}
