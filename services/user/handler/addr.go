package handler

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gomall/gateway/types/resp/common"
	user "gomall/kitex_gen/user"
	"gomall/services/user/dal/db"
	"gomall/services/user/dal/model"
	"gorm.io/gorm"
)

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
		if err = copier.Copy(address, v); err != nil {
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

	// TODO : 考虑要不要做去重判断

	address := &model.Address{ // TODO : 后面统一库表参数后 ,用copy包
		Aid:     uuid.New().String(),
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

	err = db.ModifyAddress(req.GetId(), req.GetAid(), address)
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

	// 判断记录是否存在
	r, err := db.GetAddressById(req.GetId(), req.GetId())
	if r == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		res.StatusCode = common.CodeRecordNotFound
		return
	}

	err = db.ModifyUserInfo(req.GetId(), &model.User{DefaultAddrId: req.GetAid()})
	if err != nil {
		zap.L().Error("Failed to update user info", zap.Error(err))
		return
	}

	res.StatusCode = common.CodeSuccess
	return
}
