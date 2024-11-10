package main

import (
	"context"
	"fmt"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"gomall/common/logs"
	"gomall/gateway/client"
	"gomall/services/user/config"
)

func main() {
	// 加载配置 TODO

	// 初始化日志
	logs.LogInit()

	// kitex 版链路追踪 					TODO 未测试
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.ServerName),                                                            // 配置服务名称
		provider.WithExportEndpoint(fmt.Sprintf("%s:%d", conde.Common.Jaeger.Host, config.Common.Jaeger.Port)), // Jaeger导出地址
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	// 服务发现 TODO
	r := _

	client.InitClient(r)

	// 启动路由

}
