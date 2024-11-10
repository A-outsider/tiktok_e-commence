package client

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"gomall/kitex_gen/auth/authservice"
	"gomall/kitex_gen/cart/cartservice"
	"gomall/kitex_gen/checkout/checkoutservice"
	"gomall/kitex_gen/order/orderservice"
	"gomall/kitex_gen/payment/paymentservice"
	"gomall/kitex_gen/product/productcatalogservice"
	"gomall/kitex_gen/user/userservice"
	"time"
)

var (
	UserCli     userservice.Client
	AuthCli     authservice.Client
	CartCli     cartservice.Client
	CheckoutCli checkoutservice.Client
	OrderCli    orderservice.Client
	PaymentCli  paymentservice.Client
	ProductCli  productcatalogservice.Client
)

func InitClient(resolver discovery.Resolver) {

	// TODO : 把etcd的服务发现的配置补齐

	// Initialize User Client
	UserCli = userservice.MustNewClient(
		"user",
		client.WithResolver(resolver),
		client.WithRPCTimeout(time.Second*3),
		client.WithSuite(tracing.NewClientSuite()),
	)

	// Initialize Auth Client
	AuthCli = authservice.MustNewClient(
		"auth",
		client.WithResolver(resolver),
		client.WithRPCTimeout(time.Second*3),
		client.WithSuite(tracing.NewClientSuite()),
	)

	// Initialize Cart Client
	CartCli = cartservice.MustNewClient(
		"cart",
		client.WithResolver(resolver),
		client.WithRPCTimeout(time.Second*3),
		client.WithSuite(tracing.NewClientSuite()),
	)

	// Initialize Checkout Client
	CheckoutCli = checkoutservice.MustNewClient(
		"checkout",
		client.WithResolver(resolver),
		client.WithRPCTimeout(time.Second*3),
		client.WithSuite(tracing.NewClientSuite()),
	)

	// Initialize Order Client
	OrderCli = orderservice.MustNewClient(
		"order",
		client.WithResolver(resolver),
		client.WithRPCTimeout(time.Second*3),
		client.WithSuite(tracing.NewClientSuite()),
	)

	// Initialize Payment Client
	PaymentCli = paymentservice.MustNewClient(
		"payment",
		client.WithResolver(resolver),
		client.WithRPCTimeout(time.Second*3),
		client.WithSuite(tracing.NewClientSuite()),
	)

	// Initialize Product Client
	ProductCli = productcatalogservice.MustNewClient(
		"product",
		client.WithResolver(resolver),
		client.WithRPCTimeout(time.Second*3),
		client.WithSuite(tracing.NewClientSuite()),
	)
}
