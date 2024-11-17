package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jinzhu/copier"
	"gomall/gateway/controller"
	"gomall/gateway/rpc"
	"gomall/gateway/types/req"
	"gomall/gateway/types/resp"
	"gomall/gateway/types/resp/common"
	"gomall/gateway/utils/check"
	rpcUser "gomall/kitex_gen/user"
	"gomall/kitex_gen/user/userservice"
	"path/filepath"
)

type Api struct {
	client userservice.Client
}

func NewApi() *Api {
	return &Api{
		client: rpc.GetUserClient(),
	}
}

// 获取用户信息
func (api *Api) GetUserInfo(ctx context.Context, c *app.RequestContext) {
	userId := c.GetString("userId")

	// 参数绑定
	ctrl := controller.NewCtrl[struct{}](c)

	kitexReq := &rpcUser.GetUserInfoReq{Id: userId}
	result, _ := api.client.GetUserInfo(ctx, kitexReq)

	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	respModel := new(resp.GetUserInfoResp)
	if err := copier.Copy(respModel, result); err != nil {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	ctrl.WithDataJSON(common.CodeSuccess, respModel)
}

// 修改用户信息
func (api *Api) ModifyUserInfo(ctx context.Context, c *app.RequestContext) {
	ctrl := controller.NewCtrl[req.ModifyUserInfoReq](c)

	if err := c.BindForm(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	kitexReq := &rpcUser.ModifyUserInfoReq{Id: c.GetString("userId")}
	if err := copier.Copy(kitexReq, ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	result, _ := api.client.ModifyUserInfo(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	ctrl.NoDataJSON(result.GetStatusCode())
}

// 删除用户
func (api *Api) DeleteUser(ctx context.Context, c *app.RequestContext) {
	ctrl := controller.NewCtrl[req.DeleteUserReq](c)

	if err := c.BindForm(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	userId := c.GetString("userId")

	kitexReq := &rpcUser.DeleteUserReq{Id: userId}
	if err := copier.Copy(kitexReq, ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	result, _ := api.client.DeleteUser(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	ctrl.NoDataJSON(result.GetStatusCode())
}

// 上传头像
func (api *Api) UploadAvatar(ctx context.Context, c *app.RequestContext) {

	ctrl := controller.NewCtrl[struct{}](c)

	file, err := c.FormFile("avatar")
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

	kitexReq := &rpcUser.UploadAvatarReq{
		Id:   c.GetString("userId"),
		Body: body,
		Ext:  fileExt,
	}

	result, _ := api.client.UploadAvatar(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	ctrl.NoDataJSON(result.GetStatusCode())
}

// 获取地址列表
func (api *Api) GetAddressList(ctx context.Context, c *app.RequestContext) {

	ctrl := controller.NewCtrl[struct{}](c)

	kitexReq := &rpcUser.GetAddressListReq{Id: c.GetString("userId")}

	result, _ := api.client.GetAddressList(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	if result.GetStatusCode() != common.CodeSuccess {
		ctrl.NoDataJSON(result.GetStatusCode())
		return
	}

	// TODO : 数组类型 , 全都要拷贝
	respModel := new(resp.GetAddressListResp)
	if err := copier.Copy(respModel, result); err != nil {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	ctrl.WithDataJSON(common.CodeSuccess, respModel)
}

// 添加地址
func (api *Api) AddAddress(ctx context.Context, c *app.RequestContext) {
	ctrl := controller.NewCtrl[req.AddAddressReq](c)

	if err := c.BindForm(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	kitexReq := &rpcUser.AddAddressReq{Id: c.GetString("userId")}
	if err := copier.Copy(kitexReq, ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	result, _ := api.client.AddAddress(ctx, kitexReq)

	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	ctrl.NoDataJSON(result.GetStatusCode())
}

// 修改地址
func (api *Api) ModifyAddress(ctx context.Context, c *app.RequestContext) {
	ctrl := controller.NewCtrl[req.ModifyAddressReq](c)

	if err := c.BindForm(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	kitexReq := &rpcUser.ModifyAddressReq{Id: c.GetString("userId")}
	if err := copier.Copy(kitexReq, ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	result, _ := api.client.ModifyAddress(ctx, kitexReq)
	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	ctrl.NoDataJSON(result.GetStatusCode())
}

// 删除地址
func (api *Api) DeleteAddress(ctx context.Context, c *app.RequestContext) {
	ctrl := controller.NewCtrl[req.DeleteAddressReq](c)

	if err := c.BindForm(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	kitexReq := &rpcUser.DeleteAddressReq{Id: c.GetString("userId")}
	if err := copier.Copy(kitexReq, ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	result, _ := api.client.DeleteAddress(ctx, kitexReq)

	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	ctrl.NoDataJSON(result.GetStatusCode())
}

// 设置默认地址
func (api *Api) SetDefaultAddress(ctx context.Context, c *app.RequestContext) {
	ctrl := controller.NewCtrl[req.SetDefaultAddressReq](c)

	if err := c.BindForm(ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	kitexReq := &rpcUser.SetDefaultAddressReq{Id: c.GetString("userId")}
	if err := copier.Copy(kitexReq, ctrl.Request); err != nil {
		ctrl.NoDataJSON(common.CodeInvalidParams)
		return
	}

	result, _ := api.client.SetDefaultAddress(ctx, kitexReq)

	if result == nil || result.GetStatusCode() == 0 {
		ctrl.NoDataJSON(common.CodeServerBusy)
		return
	}

	ctrl.NoDataJSON(result.GetStatusCode())
}
