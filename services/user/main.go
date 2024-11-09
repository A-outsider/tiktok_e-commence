package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/kitex-contrib/registry-etcd/retry"
	cc "gomall/common/config"
	"gomall/common/logs"
	"gomall/kitex_gen/user/userservice"
	"gomall/services/user/config"
	"gomall/services/user/handler"
	"net"
	"time"
)

func main() {

	// 配置初始化
	suite := cc.InitConfigClient(config.ServerName, config.ServerName, config.MID, config.EtcdAddr, config.GetConf())

	// 初始化日志
	logs.LogInit()

	// 服务注册
	addr, _ := net.ResolveTCPAddr("tcp", config.ServerAddr)

	retryConfig := retry.NewRetryConfig(
		retry.WithMaxAttemptTimes(10),
		retry.WithObserveDelay(20*time.Second),
		retry.WithRetryDelay(5*time.Second),
	)

	r, err := etcd.NewEtcdRegistryWithRetry([]string{config.EtcdAddr}, retryConfig) // r 不能重复使用.
	if err != nil {
		panic(err)
	}

	svr := server.NewServer(server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.ServerName}), server.WithRegistry(r), server.WithServiceAddr(addr), server.WithSuite(suite), server.WithRefuseTrafficWithoutServiceName(), server.WithMetaHandler(transmeta.ServerTTHeaderHandler))

	err = userservice.RegisterService(svr, handler.NewUserServiceImpl())

	if err != nil {
		panic(err)
	}

	err = svr.Run()

	if err != nil {
		panic(err)
	}

}
