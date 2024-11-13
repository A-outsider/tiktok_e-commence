package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/kitex-contrib/registry-etcd/retry"
	cc "gomall/common/config"
	"gomall/common/database"
	"gomall/common/logs"
	"gomall/kitex_gen/user/userservice"
	"gomall/services/user/config"
	"gomall/services/user/dal/db"
	"gomall/services/user/handler"
	"net"
	"time"
)

func main() {

	// 配置初始化
	etcdSuite := cc.InitConfigClient(config.ServerName, config.ServerName, config.MID, config.EtcdAddr, config.GetConf())
	// 初始化日志
	logs.LogInit(config.ServerName)

	// otel 版链路追踪
	//shutdown, err := database.InitTracer(config.ServerName)
	//if err != nil {
	//	log.Panic(err)
	//}
	//defer func() {
	//	if err = shutdown(context.Background()); err != nil {
	//		log.Panicf("Error shutting down tracer provider: %v\n", err)
	//	}
	//}()

	// kitex 版链路追踪 					TODO 未测试
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.ServerName), // 配置服务名称
		provider.WithExportEndpoint(fmt.Sprintf("%s:%d", config.GetConf().Jaeger.Host, config.GetConf().Jaeger.Port)), // Jaeger导出地址
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	// 连MySQL
	if err := db.InitDb(config.GetConf().Mysql); err != nil {
		panic(err)
	}

	// 连 Redis
	database.InitRedis(config.GetConf().Redis)

	// 服务注册
	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", config.GetConf().Service.Host, config.GetConf().Service.Port))

	retryConfig := retry.NewRetryConfig( // 重试策略
		retry.WithMaxAttemptTimes(10),
		retry.WithObserveDelay(20*time.Second),
		retry.WithRetryDelay(5*time.Second),
	)

	r, err := etcd.NewEtcdRegistryWithRetry([]string{config.EtcdAddr}, retryConfig) // r 不能重复使用.
	if err != nil {
		panic(err)
	}

	svr := server.NewServer(
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.ServerName}), // 服务基本信息
		server.WithServiceAddr(addr),                            // 服务地址
		server.WithRegistry(r),                                  // 服务注册中心
		server.WithRefuseTrafficWithoutServiceName(),            // 拒绝没有服务名的请求
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler), // 元数据处理器
		server.WithSuite(etcdSuite),                             // etcd套件
		server.WithSuite(tracing.NewServerSuite()),              // opentelemetry 套件
	)

	if err = userservice.RegisterService(svr, handler.NewUserServiceImpl()); err != nil {
		panic(err)
	}
	if err = svr.Run(); err != nil {
		panic(err)
	}

}
