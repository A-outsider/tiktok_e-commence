package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/kitex-contrib/registry-etcd/retry"
	cc "gomall/common/config"
	"gomall/common/logs"
	auth "gomall/kitex_gen/auth/authservice"
	"gomall/services/auth/config"
	"gomall/services/auth/handler"
	"gomall/services/auth/initialize"
	"net"
	"time"
)

func main() {

	// 配置初始化
	etcdSuite := cc.InitConfigClient(config.ServerName, config.ServerName, config.MID, config.EtcdAddr, config.GetConf())
	// 初始化日志
	logs.LogInit(config.ServerName)

	// kitex 版链路追踪
	//p := provider.NewOpenTelemetryProvider(
	//	provider.WithServiceName(config.ServerName), // 配置服务名称
	//	provider.WithExportEndpoint(fmt.Sprintf("%s:%d", config.GetConf().Jaeger.Host, config.GetConf().Jaeger.Port)), // Jaeger导出地址
	//	provider.WithInsecure(),
	//	provider.WithEnableMetrics(false),
	//)
	//defer p.Shutdown(context.Background())

	// 初始化一系列主件
	initialize.Init()

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
		//server.WithSuite(tracing.NewServerSuite()),              // opentelemetry 套件
	)

	if err = auth.RegisterService(svr, handler.NewAuthServiceImpl()); err != nil {
		panic(err)
	}
	if err = svr.Run(); err != nil {
		panic(err)
	}

}
