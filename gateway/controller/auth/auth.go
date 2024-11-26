package auth

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jinzhu/copier"
	"gomall/gateway/controller"
	"gomall/gateway/rpc"
	"gomall/gateway/types/req"
	"gomall/gateway/types/resp"
	"gomall/gateway/types/resp/common"
	rpcAuth "gomall/kitex_gen/auth"
	"gomall/kitex_gen/auth/authservice"
)

type Api struct {
	client authservice.Client
}

func NewApi() *Api {
	return &Api{
		client: rpc.GetAuthClient(),
	}
}

// 登录 - 验证码
func (api *Api) LoginByCode(ctx context.Context, c *app.RequestContext) {

	// 参数绑定
	ctrl := controller.NewCtrl[req.LoginByCodeReq](c)
	if err := c.BindForm(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 转模型
	kitexReq := new(rpcAuth.LoginByCodeReq)
	err := copier.Copy(kitexReq, ctrl.Request)
	if err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 调用 RPC 方法
	result, _ := api.client.LoginByCode(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	// 转模型
	respModel := new(resp.LoginByCodeResp)
	err = copier.Copy(respModel, result)
	if err != nil {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	ctrl.WithDataJSON(common.CodeSuccess, respModel)
}

// 登录 - 密码
func (api *Api) LoginByPwd(ctx context.Context, c *app.RequestContext) {
	ctrl := controller.NewCtrl[req.LoginByPwdReq](c)

	if err := c.BindForm(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 转模型
	kitexReq := new(rpcAuth.LoginByPwdReq)
	if err := copier.Copy(kitexReq, ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	result, _ := api.client.LoginByPwd(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	// 转模型
	respModel := new(resp.LoginByPwdResp)
	if err := copier.Copy(respModel, result); err != nil {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	ctrl.WithDataJSON(common.CodeSuccess, respModel)
}

// 注册
func (api *Api) Register(ctx context.Context, c *app.RequestContext) {
	ctrl := controller.NewCtrl[req.RegisterReq](c)

	if err := c.BindForm(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 转模型
	kitexReq := new(rpcAuth.RegisterReq)
	if err := copier.Copy(kitexReq, ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	result, _ := api.client.Register(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	ctrl.NoDataJSON(common.CodeSuccess)
}

// 发送验证码
func (api *Api) SendPhoneCode(ctx context.Context, c *app.RequestContext) {
	ctrl := controller.NewCtrl[req.SendPhoneCodeReq](c)

	if err := c.BindForm(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 转模型
	kitexReq := new(rpcAuth.SendPhoneCodeReq)
	if err := copier.Copy(kitexReq, ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	result, _ := api.client.SendPhoneCode(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	ctrl.NoDataJSON(common.CodeSuccess)
}

// 发送邮箱验证码
func (api *Api) SendEmailCode(ctx context.Context, c *app.RequestContext) {
	ctrl := controller.NewCtrl[req.SendEmailCodeReq](c)

	if err := c.BindForm(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 转模型
	kitexReq := new(rpcAuth.SendEmailCodeReq)
	if err := copier.Copy(kitexReq, ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	result, _ := api.client.SendEmailCode(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	ctrl.NoDataJSON(common.CodeSuccess)
}

// 刷新 Token
func (api *Api) RefreshToken(ctx context.Context, c *app.RequestContext) {
	ctrl := controller.NewCtrl[req.RefreshTokenReq](c)

	if err := c.BindForm(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	// 转模型
	kitexReq := new(rpcAuth.RefreshTokenReq)
	if err := copier.Copy(kitexReq, ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	result, _ := api.client.RefreshToken(ctx, kitexReq)

	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	// 转模型
	respModel := new(resp.RefreshTokenResp)
	if err := copier.Copy(respModel, result); err != nil {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	ctrl.WithDataJSON(common.CodeSuccess, respModel)
}

// 显示图片验证码
func (api *Api) ShowPhotoCaptcha(ctx context.Context, c *app.RequestContext) {
	ctrl := controller.NewCtrl[req.None](c)

	result, _ := api.client.ShowPhotoCaptcha(ctx, nil)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	// 转模型
	respModel := new(resp.ShowPhotoCaptchaResp)
	if err := copier.Copy(respModel, result); err != nil {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	ctrl.WithDataJSON(common.CodeSuccess, respModel)
}
