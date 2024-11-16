package handler

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gomall/gateway/types/resp/common"
	user "gomall/kitex_gen/user"
	"gomall/services/user/config"
	"gomall/services/user/dal/cache"
	"gomall/services/user/dal/db"
	"gomall/services/user/dal/model"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

type UserServiceImpl struct{}

// NewUserServiceImpl creates a new instance of UserServiceImpl.
func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{}
}

// GetUserInfo implements the UserService interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (res *user.GetUserInfoResp, err error) {
	res = new(user.GetUserInfoResp)
	res.StatusCode = common.CodeServerBusy

	userInfo, err := db.GetUserByID(req.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.StatusCode = common.CodeUserNotExist
			return
		}
		zap.L().Error("Failed to fetch user info", zap.Error(err))
		return
	}

	// 模型转换
	if err = copier.Copy(res, userInfo); err != nil {
		zap.L().Error("Failed to copy user info", zap.Error(err))
		return
	}

	res.StatusCode = common.CodeSuccess
	return
}

// ModifyUserInfo implements the UserService interface.
func (s *UserServiceImpl) ModifyUserInfo(ctx context.Context, req *user.ModifyUserInfoReq) (res *user.ModifyUserInfoResp, err error) {
	res = new(user.ModifyUserInfoResp)
	res.StatusCode = common.CodeServerBusy

	u := new(model.User)
	if err = copier.Copy(u, req); err != nil {
		zap.L().Error("Failed to copy user info", zap.Error(err))
		return
	}

	err = db.ModifyUserInfo(req.GetId(), u)
	if err != nil {
		zap.L().Error("Failed to update user info", zap.Error(err))
		return
	}

	res.StatusCode = common.CodeSuccess
	return
}

// DeleteUser implements the UserService interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserReq) (res *user.DeleteUserResp, err error) {
	res = new(user.DeleteUserResp)
	res.StatusCode = common.CodeServerBusy

	// 校验手机验证码
	code, err := cache.Get(cache.GetPhoneCodeKey(req.GetPhone()))
	if err != nil {
		return nil, err
	}

	if code != req.GetAuthCode() {
		res.StatusCode = common.CodeInvalidCaptcha
		return
	}

	err = db.DeleteUserById(req.GetId())
	if err != nil {
		zap.L().Error("Failed to delete user", zap.Error(err))
		return
	}

	res.StatusCode = common.CodeSuccess
	return
}

// UploadAvatar implements the UserService interface.
func (s *UserServiceImpl) UploadAvatar(ctx context.Context, req *user.UploadAvatarReq) (res *user.UploadAvatarResp, err error) {
	res = new(user.UploadAvatarResp)
	res.StatusCode = common.CodeServerBusy

	// 将原头像删除（如果存在的话）
	u, err := db.GetUser(req.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.StatusCode = common.CodeUserNotExist
			return
		}
		zap.L().Error("Failed to get user", zap.Error(err))
		return
	}

	if len(u.AvatarPath) != 0 {
		_ = os.Remove(filepath.Join(config.GetConf().Static.AvatarPath, u.AvatarPath))
	}

	// 保存新的头像
	fileName := uuid.New().String() + req.GetExt()
	os.MkdirAll(config.GetConf().Static.AvatarPath, 0755)
	if err = os.WriteFile(filepath.Join(config.GetConf().Static.AvatarPath, fileName), req.GetBody(), 0644); err != nil {
		zap.L().Error("Failed to save avatar", zap.Error(err))
		return
	}

	// 更新数据库
	if err = db.ModifyUserInfo(req.GetId(), &model.User{AvatarPath: fileName}); err != nil {
		zap.L().Error("Failed to update user info", zap.Error(err))
		return
	}

	res.StatusCode = common.CodeSuccess
	return
}

// GetAddressList implements the UserService interface.
func (s *UserServiceImpl) GetAddressList(ctx context.Context, req *user.GetAddressListReq) (res *user.GetAddressListResp, err error) {
	res = new(user.GetAddressListResp)
	res.StatusCode = common.CodeServerBusy

	var list []*model.Address
	list, err = db.GetAddressList(req.GetId())
	if err != nil {
		zap.L().Error("Failed to get address list", zap.Error(err))
		return
	}

	// 模型转化
	for _, v := range list {
		address := new(user.Address)
		if err = copier.Copy(&address, v); err != nil {
			zap.L().Error("Failed to copy address", zap.Error(err))
			return
		}
		res.Addresses = append(res.Addresses, address)
	}

	res.StatusCode = common.CodeSuccess
	return
}

// AddAddress implements the UserService interface.
func (s *UserServiceImpl) AddAddress(ctx context.Context, req *user.AddAddressReq) (res *user.AddAddressResp, err error) {
	res = new(user.AddAddressResp)
	res.StatusCode = common.CodeServerBusy

	address := &model.Address{ // TODO : 后面统一库表参数后 ,用copy包
		Uid:     req.GetId(),
		Name:    req.GetName(),
		Phone:   req.GetPhone(),
		Address: req.GetAddress(),
	}

	err = db.CreateAddress(address)
	if err != nil {
		zap.L().Error("Failed to create address", zap.Error(err))
		return
	}

	res.StatusCode = common.CodeSuccess
	return
}

// ModifyAddress implements the UserService interface.
func (s *UserServiceImpl) ModifyAddress(ctx context.Context, req *user.ModifyAddressReq) (res *user.ModifyAddressResp, err error) {
	res = new(user.ModifyAddressResp)
	res.StatusCode = common.CodeServerBusy

	address := &model.Address{
		Name:    req.GetName(),
		Phone:   req.GetPhone(),
		Address: req.GetAddress(),
	}

	err = db.ModifyAddress(req.GetId(), address)
	if err != nil {
		zap.L().Error("Failed to update address", zap.Error(err))
		return
	}

	res.StatusCode = common.CodeSuccess
	return
}

// DeleteAddress implements the UserService interface.
func (s *UserServiceImpl) DeleteAddress(ctx context.Context, req *user.DeleteAddressReq) (res *user.DeleteAddressResp, err error) {
	res = new(user.DeleteAddressResp)
	res.StatusCode = common.CodeServerBusy

	ok := db.DeleteAddress(req.GetId(), req.GetAid())
	if !ok {
		res.StatusCode = common.CodeRecordNotFound
		return
	}

	res.StatusCode = common.CodeSuccess
	return
}

// SetDefaultAddress implements the UserService interface.
func (s *UserServiceImpl) SetDefaultAddress(ctx context.Context, req *user.SetDefaultAddressReq) (res *user.SetDefaultAddressResp, err error) {
	res = new(user.SetDefaultAddressResp)
	res.StatusCode = common.CodeServerBusy

	err = db.ModifyUserInfo(req.GetId(), &model.User{DefaultId: req.GetAid()})
	if err != nil {
		zap.L().Error("Failed to update user info", zap.Error(err))
		return
	}

	res.StatusCode = common.CodeSuccess
	return
}
