package order

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jinzhu/copier"
	"gomall/gateway/controller"
	"gomall/gateway/rpc"
	"gomall/gateway/types/req"
	"gomall/gateway/types/resp/common"
	rpcOrder "gomall/kitex_gen/order"
	order "gomall/kitex_gen/order/orderservice"
)

type Api struct {
	client order.Client
}

func NewApi() *Api {
	return &Api{
		client: rpc.GetOrderClient(),
	}
}

func (api *Api) PlaceOrder(ctx context.Context, c *app.RequestContext) {
	// 参数绑定
	ctrl := controller.NewCtrl[req.PlaceOrderReq](c)
	if err := c.Bind(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 转模型
	kitexReq := new(rpcOrder.PlaceOrderReq)
	err := copier.Copy(kitexReq, ctrl.Request)
	if err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}
	kitexReq.UserId = c.GetString("userId")

	// 调用 RPC 方法
	result, _ := api.client.PlaceOrder(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	ctrl.WithDataJSON(result.GetStatusCode(), result.Order)
}

func (api *Api) ListOrder(ctx context.Context, c *app.RequestContext) {
	// 参数绑定
	ctrl := controller.NewCtrl[req.None](c)

	// 转模型
	kitexReq := new(rpcOrder.ListOrderReq)
	kitexReq.UserId = c.GetString("userId")

	// 调用 RPC 方法
	result, _ := api.client.ListOrder(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	ctrl.WithDataJSON(result.GetStatusCode(), result.Orders)
}

func (api *Api) ListOrderFromSeller(ctx context.Context, c *app.RequestContext) {
	// 参数绑定
	ctrl := controller.NewCtrl[req.None](c)

	// 转模型
	kitexReq := new(rpcOrder.ListOrderFromSellerReq)
	kitexReq.SellerId = c.GetString("userId")

	// 调用 RPC 方法
	result, _ := api.client.ListOrderFromSeller(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	ctrl.WithDataJSON(result.GetStatusCode(), result.Orders)
}

func (api *Api) MarkOrderShipped(ctx context.Context, c *app.RequestContext) {
	// 参数绑定
	ctrl := controller.NewCtrl[req.ChangeStatusReq](c)

	// 转模型
	kitexReq := new(rpcOrder.MarkOrderShippedReq)
	kitexReq.UserId = c.GetString("userId")

	// 调用 RPC 方法
	result, _ := api.client.MarkOrderShipped(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	ctrl.NoDataJSON(result.StatusCode)
}

func (api *Api) MarkOrderCompleted(ctx context.Context, c *app.RequestContext) {
	// 参数绑定
	ctrl := controller.NewCtrl[req.ChangeStatusReq](c)

	// 转模型
	kitexReq := new(rpcOrder.MarkOrderCompletedReq)
	kitexReq.UserId = c.GetString("userId")

	// 调用 RPC 方法
	result, _ := api.client.MarkOrderCompleted(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	ctrl.NoDataJSON(result.StatusCode)
}
