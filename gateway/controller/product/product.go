package product

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jinzhu/copier"
	"gomall/gateway/controller"
	"gomall/gateway/rpc"
	"gomall/gateway/types/req"
	"gomall/gateway/types/resp/common"
	rpcProduct "gomall/kitex_gen/product"
	"gomall/kitex_gen/product/productcatalogservice"
)

type Api struct {
	client productcatalogservice.Client
}

func NewApi() *Api {
	return &Api{
		client: rpc.GetProductClient(),
	}
}

func (api *Api) AddProduct(ctx context.Context, c *app.RequestContext) {
	// 参数绑定
	ctrl := controller.NewCtrl[req.AddProductReq](c)
	if err := c.BindForm(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 转模型
	kitexReq := new(rpcProduct.AddProductReq)
	kitexReq.Product = new(rpcProduct.Product)
	err := copier.Copy(kitexReq.Product, ctrl.Request)
	if err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}
	kitexReq.Product.Uid = c.GetString("userId")

	// 调用 RPC 方法
	result, _ := api.client.AddProduct(ctx, kitexReq)
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

func (api *Api) ListProducts(ctx context.Context, c *app.RequestContext) {
	// 参数绑定
	ctrl := controller.NewCtrl[req.SearchProductByCategoryReq](c)
	if err := c.BindForm(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 转模型
	kitexReq := new(rpcProduct.ListProductsReq)
	err := copier.Copy(kitexReq, ctrl.Request)
	if err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 调用 RPC 方法
	result, _ := api.client.ListProducts(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	ctrl.WithDataJSON(result.GetStatusCode(), result.Products)
}
