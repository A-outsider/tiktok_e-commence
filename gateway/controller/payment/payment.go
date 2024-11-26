package payment

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jinzhu/copier"
	"gomall/gateway/controller"
	"gomall/gateway/rpc"
	"gomall/gateway/types/req"
	"gomall/gateway/types/resp/common"
	rpcPayment "gomall/kitex_gen/payment"
	"gomall/kitex_gen/payment/paymentservice"
	"net/http"
)

type Api struct {
	client paymentservice.Client
}

func NewApi() *Api {
	return &Api{
		client: rpc.GetPaymentClient(),
	}
}

func (api *Api) CreatePayment(ctx context.Context, c *app.RequestContext) {
	ctrl := controller.NewCtrl[req.CreatePaymentReq](c)
	if err := c.BindForm(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	kitexReq := new(rpcPayment.CreatePaymentReq)
	if err := copier.Copy(kitexReq, ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	kitexReq.UserId = c.GetString("userId")

	result, _ := api.client.CreatePayment(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	c.JSON(http.StatusOK, result.GetPaymentUrl()) // 用于测试
	//c.Redirect(http.StatusTemporaryRedirect, []byte(result.GetPaymentUrl()))
}

func (api *Api) PayCallback(ctx context.Context, c *app.RequestContext) {
	ctrl := controller.NewCtrl[req.None](c)

	rawData := c.Request.URI().QueryString()

	kitexReq := new(rpcPayment.PayCallbackReq)
	kitexReq.RawData = rawData

	result, _ := api.client.PayCallback(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	// TODO 支付成功 , 跳转到成功页面
	c.JSON(http.StatusOK, "success")
	//c.Redirect(http.StatusTemporaryRedirect, []byte("..."))
}

func (api *Api) PayNotify(ctx context.Context, c *app.RequestContext) {
	ctrl := controller.NewCtrl[req.None](c)

	rawData, err := c.Body()
	if err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	kitexReq := new(rpcPayment.PayNotifyReq)
	kitexReq.RawData = rawData

	result, _ := api.client.PayNotify(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	c.String(http.StatusOK, "success") // 模拟 ACKNotification 方法来响应支付宝服务器
}
