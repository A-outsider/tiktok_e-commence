package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"gomall/gateway/middleware"
	"gomall/gateway/router/model"
)

func Register(r *server.Hertz) {
	r.Use(middleware.CORS())
	v1 := r.Group("/api/v1")

	// 测试
	v1.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	// 分模块
	model.RegisterAuth(v1.Group("/auth"))
	model.RegisterUser(v1.Group("/user"))

}
