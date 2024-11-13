package rpc

//import (
//	"github.com/cloudwego/kitex/client"
//	"github.com/cloudwego/kitex/pkg/transmeta"
//	"github.com/cloudwego/kitex/transport"
//	"github.com/kitex-contrib/obs-opentelemetry/tracing"
//	etcd "github.com/kitex-contrib/registry-etcd"
//	"gomall/common/config"
//	"gomall/kitex_gen/order/orderservice"
//)
//
//var orderCli orderservice.Client
//
//func initOrder() {
//	r, err := etcd.NewEtcdResolver([]string{config.EtcdAddr})
//	if err != nil {
//		panic(err)
//	}
//	orderCli, err = orderservice.NewClient(
//		"order",
//		client.WithResolver(r),
//		client.WithTransportProtocol(transport.TTHeader),
//		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
//		client.WithSuite(tracing.NewClientSuite()),
//	)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func GetOrderClient() orderservice.Client {
//	return orderCli
//}
