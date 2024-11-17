package main

import (
	cc "gomall/common/config"
	"gomall/common/logs"
	"gomall/gateway/config"
	"gomall/gateway/router"
	"gomall/gateway/rpc"
	"gomall/gateway/utils/role"
)

func main() {
	// 加载配置
	cc.InitConfigClient(config.ServerName, config.ServerName, config.MID, config.EtcdAddr, config.GetConf())

	// 初始化日志
	logs.LogInit(config.ServerName)

	// 初始化策略模型
	role.InitCasbin()

	// kitex 版链路追踪 					TODO 未测试
	//p := provider.NewOpenTelemetryProvider(
	//	provider.WithServiceName(config.ServerName), // 配置服务名称
	//	provider.WithExportEndpoint(fmt.Sprintf("%s:%d", config.GetConf().Jaeger.Host, config.GetConf().Jaeger.Port)), // Jaeger导出地址
	//	provider.WithInsecure(),
	//	provider.WithEnableMetrics(false),
	//)
	//defer p.Shutdown(context.Background())

	// 服务发现
	rpc.Init()

	// 启动路由
	router.InitRouter().Spin()
}
