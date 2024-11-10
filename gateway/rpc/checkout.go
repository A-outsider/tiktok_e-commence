package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	etcd "github.com/kitex-contrib/registry-etcd"
	"gomall/common/config"
	"gomall/kitex_gen/checkout/checkoutservice"
)

var checkoutCli checkoutservice.Client

func initCheckout() {
	r, err := etcd.NewEtcdResolver([]string{config.EtcdAddr})
	if err != nil {
		panic(err)
	}
	checkoutCli, err = checkoutservice.NewClient("checkout", client.WithResolver(r), client.WithTransportProtocol(transport.TTHeader), client.WithMetaHandler(transmeta.ClientTTHeaderHandler))
	if err != nil {
		panic(err)
	}
}

func GetCheckoutClient() checkoutservice.Client {
	return checkoutCli
}
