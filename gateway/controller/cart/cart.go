package cart

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jinzhu/copier"
	"gomall/gateway/controller"
	"gomall/gateway/rpc"
	"gomall/gateway/types/req"
	"gomall/gateway/types/resp/common"
	rpcCart "gomall/kitex_gen/cart"
	"gomall/kitex_gen/cart/cartservice"
)

type Api struct {
	client cartservice.Client
}

func NewApi() *Api {
	return &Api{
		client: rpc.GetCartClient(),
	}
}

func (api *Api) AddItem(ctx context.Context, c *app.RequestContext) {
	// 参数绑定
	ctrl := controller.NewCtrl[req.AddItemReq](c)
	if err := c.BindForm(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 转模型
	kitexReq := new(rpcCart.AddItemReq)
	kitexReq.Item = new(rpcCart.CartItem)
	err := copier.Copy(kitexReq.Item, ctrl.Request)
	if err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}
	kitexReq.UserId = c.GetString("userId")

	// 调用 RPC 方法
	result, _ := api.client.AddItem(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	ctrl.NoDataJSON(result.GetStatusCode())
}

func (api *Api) GetCart(ctx context.Context, c *app.RequestContext) {

	ctrl := controller.NewCtrl[req.None](c)
	// 转模型
	kitexReq := new(rpcCart.GetCartReq)
	kitexReq.UserId = c.GetString("userId")

	// 调用 RPC 方法
	result, _ := api.client.GetCart(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	ctrl.WithDataJSON(result.GetStatusCode(), result.Cart)
}

func (api *Api) EmptyCart(ctx context.Context, c *app.RequestContext) {
	ctrl := controller.NewCtrl[req.None](c)

	kitexReq := new(rpcCart.EmptyCartReq)
	kitexReq.UserId = c.GetString("userId")

	result, _ := api.client.EmptyCart(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}
	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	ctrl.NoDataJSON(result.GetStatusCode())
}
