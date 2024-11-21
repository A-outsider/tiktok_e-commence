package handler

import (
	"context"
	"go.uber.org/zap"
	"gomall/gateway/types/resp/common"
	cart "gomall/kitex_gen/cart"
	"gomall/services/cart/dal/cache"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// AddItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddItem(ctx context.Context, req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// TODO: Your code here...
	resp = new(cart.AddItemResp)
	resp.StatusCode = common.CodeServerBusy

	err = cache.AddItem(ctx, req.UserId, req.Item.ProductId, req.Item.Quantity)
	if err != nil {
		zap.L().Error("add item failed", zap.Error(err))
		return
	}

	resp.StatusCode = common.CodeSuccess
	return
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// TODO: Your code here...
	resp = new(cart.GetCartResp)
	resp.StatusCode = common.CodeServerBusy
	resp.Cart = new(cart.Cart)

	items, err := cache.GetItems(ctx, req.GetUserId())
	if err != nil {
		zap.L().Error("get cart failed", zap.Error(err))
		return
	}

	resp.Cart.Items = items
	resp.Cart.UserId = req.UserId
	resp.StatusCode = common.CodeSuccess
	return
}

// EmptyCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// TODO: Your code here...
	resp = new(cart.EmptyCartResp)
	resp.StatusCode = common.CodeServerBusy

	err = cache.DeleteItems(ctx, req.UserId)
	if err != nil {
		zap.L().Error("empty cart failed", zap.Error(err))
		return
	}

	resp.StatusCode = common.CodeSuccess
	return
}

func NewCartServiceImpl() *CartServiceImpl {
	return &CartServiceImpl{}
}
