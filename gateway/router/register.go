package router

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hconfig "github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/logger/accesslog"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"gomall/gateway/config"
	"gomall/gateway/middleware"
	"gomall/gateway/router/model"
	"log"
)

func InitRouter() *server.Hertz {

	// 设置配置
	opts := []hconfig.Option{server.WithHostPorts(fmt.Sprintf("%s:%d", config.GetConf().Service.Host, config.GetConf().Service.Port))}

	// 初始化链路追踪配置
	tracer, cfg := hertztracing.NewServerTracer()
	opts = append(opts, tracer)

	// 集成日志系统	TODO : 待测试
	hlog.SetLevel(hlog.LevelDebug) // 可根据需求调整

	h := server.Default(opts...)
	h.Use(hertztracing.ServerMiddleware(cfg), accesslog.New()) // 设置链路追踪和日志

	registerRouter(h)

	log.Printf("server run at %s:%d\n", config.GetConf().Service.Host, config.GetConf().Service.Port)
	return h
}

func registerRouter(r *server.Hertz) {
	r.Use(middleware.CORS())

	// 获取静态文件
	r.Static("/static", config.GetConf().Static.AvatarPath) // 用户头像文件夹	TODO : 这个框架的映射好像有bug ?

	v1 := r.Group("/api/v1")

	// 测试
	v1.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	// 分模块
	model.RegisterAuth(v1.Group("/auth"))
	model.RegisterUser(v1.Group("/user"))
	model.RegisterProduct(v1.Group("/product"))
	model.RegisterCart(v1.Group("/cart"))
	model.RegisterOrder(v1.Group("/order"))
	model.RegisterPayment(v1.Group("/payment"))
}
