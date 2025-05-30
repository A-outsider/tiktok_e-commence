// Code generated by Kitex v0.10.3. DO NOT EDIT.

package orderservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	order "gomall/kitex_gen/order"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	PlaceOrder(ctx context.Context, Req *order.PlaceOrderReq, callOptions ...callopt.Option) (r *order.PlaceOrderResp, err error)
	ListOrder(ctx context.Context, Req *order.ListOrderReq, callOptions ...callopt.Option) (r *order.ListOrderResp, err error)
	ListOrderFromSeller(ctx context.Context, Req *order.ListOrderFromSellerReq, callOptions ...callopt.Option) (r *order.ListOrderFromSellerResp, err error)
	MarkOrderPaid(ctx context.Context, Req *order.MarkOrderPaidReq, callOptions ...callopt.Option) (r *order.MarkOrderPaidResp, err error)
	MakeSureOrderExpired(ctx context.Context, Req *order.MakeSureOrderExpiredReq, callOptions ...callopt.Option) (r *order.MakeSureOrderExpiredResp, err error)
	MarkOrderShipped(ctx context.Context, Req *order.MarkOrderShippedReq, callOptions ...callopt.Option) (r *order.MarkOrderShippedResp, err error)
	MarkOrderCompleted(ctx context.Context, Req *order.MarkOrderCompletedReq, callOptions ...callopt.Option) (r *order.MarkOrderCompletedResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kOrderServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kOrderServiceClient struct {
	*kClient
}

func (p *kOrderServiceClient) PlaceOrder(ctx context.Context, Req *order.PlaceOrderReq, callOptions ...callopt.Option) (r *order.PlaceOrderResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PlaceOrder(ctx, Req)
}

func (p *kOrderServiceClient) ListOrder(ctx context.Context, Req *order.ListOrderReq, callOptions ...callopt.Option) (r *order.ListOrderResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListOrder(ctx, Req)
}

func (p *kOrderServiceClient) ListOrderFromSeller(ctx context.Context, Req *order.ListOrderFromSellerReq, callOptions ...callopt.Option) (r *order.ListOrderFromSellerResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListOrderFromSeller(ctx, Req)
}

func (p *kOrderServiceClient) MarkOrderPaid(ctx context.Context, Req *order.MarkOrderPaidReq, callOptions ...callopt.Option) (r *order.MarkOrderPaidResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MarkOrderPaid(ctx, Req)
}

func (p *kOrderServiceClient) MakeSureOrderExpired(ctx context.Context, Req *order.MakeSureOrderExpiredReq, callOptions ...callopt.Option) (r *order.MakeSureOrderExpiredResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MakeSureOrderExpired(ctx, Req)
}

func (p *kOrderServiceClient) MarkOrderShipped(ctx context.Context, Req *order.MarkOrderShippedReq, callOptions ...callopt.Option) (r *order.MarkOrderShippedResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MarkOrderShipped(ctx, Req)
}

func (p *kOrderServiceClient) MarkOrderCompleted(ctx context.Context, Req *order.MarkOrderCompletedReq, callOptions ...callopt.Option) (r *order.MarkOrderCompletedResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MarkOrderCompleted(ctx, Req)
}
