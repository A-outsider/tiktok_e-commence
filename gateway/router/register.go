package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hconfig "github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"gomall/common/logs"
	"gomall/gateway/config"
	"gomall/gateway/middleware"
	"gomall/gateway/router/model"
)

func InitRouter() *server.Hertz {

	// 设置配置
	opts := []hconfig.Option{server.WithHostPorts(config.GetConf().Server.Addr)}

	// 初始化链路追踪配置
	tracer, cfg := hertztracing.NewServerTracer()
	opts = append(opts, tracer)

	// 集成日志系统	TODO : 待测试
	hlog.SetLevel(hlog.LevelDebug) // 可根据需求调整

	h := server.Default(opts...)
	h.Use(hertztracing.ServerMiddleware(cfg), logs.AccessLog()) // 设置链路追踪和日志

	registerRouter(h)

	return h
}

func registerRouter(r *server.Hertz) {
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
