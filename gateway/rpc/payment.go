package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	etcd "github.com/kitex-contrib/registry-etcd"
	"gomall/common/config"
	"gomall/kitex_gen/payment/paymentservice"
)

var paymentCli paymentservice.Client

func initPayment() {
	r, err := etcd.NewEtcdResolver([]string{config.EtcdAddr})
	if err != nil {
		panic(err)
	}
	paymentCli, err = paymentservice.NewClient("payment", client.WithResolver(r), client.WithTransportProtocol(transport.TTHeader), client.WithMetaHandler(transmeta.ClientTTHeaderHandler))
	if err != nil {
		panic(err)
	}
}

func GetPaymentClient() paymentservice.Client {
	return paymentCli
}
