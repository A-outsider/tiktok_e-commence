package rpc

import (
	"github.com/cloudwego/kitex/client"
	kClient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"gomall/common/config"
	"gomall/kitex_gen/auth/authservice"
)

var authCli authservice.Client

func initAuth() {
	r, err := etcd.NewEtcdResolver([]string{config.EtcdAddr})
	if err != nil {
		panic(err)
	}
	authCli, err = authservice.NewClient(
		"auth",
		client.WithResolver(r),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithSuite(tracing.NewClientSuite()),
		kClient.WithTracer(prometheus.NewClientTracer(":9091", "/kitexclient")), // 本地启动要做内网穿透
	)
	if err != nil {
		panic(err)
	}
}

func GetAuthClient() authservice.Client {
	return authCli
}
