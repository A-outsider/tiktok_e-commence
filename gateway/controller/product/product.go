package product

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jinzhu/copier"
	"gomall/gateway/controller"
	"gomall/gateway/rpc"
	"gomall/gateway/types/req"
	"gomall/gateway/types/resp/common"
	"gomall/gateway/utils/check"
	rpcProduct "gomall/kitex_gen/product"
	"gomall/kitex_gen/product/productcatalogservice"
	"path/filepath"
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

	file, err := c.FormFile("picture")
	if err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 校验文件格式
	fileExt := filepath.Ext(file.Filename)
	if !check.PhotoSize(file.Size) || !check.PhotoType(fileExt) {
		ctrl.NoDataJSON(common.CodeInvalidParams) // TODO 参数细化
		return
	}

	fileReader, err := file.Open()
	if err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}
	defer fileReader.Close()

	body := make([]byte, file.Size)
	_, err = fileReader.Read(body)
	if err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 转模型
	kitexReq := new(rpcProduct.AddProductReq)
	kitexReq.Product = new(rpcProduct.Product)
	err = copier.Copy(kitexReq.Product, ctrl.Request)
	if err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}
	kitexReq.Product.Uid = c.GetString("userId")
	kitexReq.Ext = fileExt
	kitexReq.Body = body

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
	if err := c.BindQuery(ctrl.Request); err != nil {
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

func (api *Api) SearchProducts(ctx context.Context, c *app.RequestContext) {
	// 参数绑定
	ctrl := controller.NewCtrl[req.SearchProductByQueryReq](c)
	if err := c.BindQuery(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 转模型
	kitexReq := new(rpcProduct.SearchProductsReq)
	err := copier.Copy(kitexReq, ctrl.Request)
	if err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 调用 RPC 方法
	result, _ := api.client.SearchProducts(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	ctrl.WithDataJSON(result.GetStatusCode(), result.Results)
}

func (api *Api) DeleteProduct(ctx context.Context, c *app.RequestContext) {
	// 参数绑定
	ctrl := controller.NewCtrl[req.DeleteProductReq](c)
	if err := c.BindForm(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 转模型
	kitexReq := new(rpcProduct.DeleteProductReq)
	err := copier.Copy(kitexReq, ctrl.Request)
	if err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 调用 RPC 方法
	result, _ := api.client.DeleteProduct(ctx, kitexReq)
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

func (api *Api) GetProduct(ctx context.Context, c *app.RequestContext) {
	// 参数绑定
	ctrl := controller.NewCtrl[req.GetProductReq](c)
	if err := c.BindQuery(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 转模型
	kitexReq := new(rpcProduct.GetProductReq)
	err := copier.Copy(kitexReq, ctrl.Request)
	if err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 调用 RPC 方法
	result, _ := api.client.GetProduct(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	ctrl.WithDataJSON(result.GetStatusCode(), result.GetProduct())
}

func (api *Api) GetRankings(ctx context.Context, c *app.RequestContext) {
	// 参数绑定
	ctrl := controller.NewCtrl[req.None](c)
	if err := c.BindQuery(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 转模型
	kitexReq := new(rpcProduct.GetRankingsReq)

	// 调用 RPC 方法
	result, _ := api.client.GetRankings(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	ctrl.WithDataJSON(result.GetStatusCode(), result.GetProductItems())
}
