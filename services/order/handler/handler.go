package handler

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gomall/gateway/types/resp/common"
	order "gomall/kitex_gen/order"
	"gomall/services/order/dal/cache"
	"gomall/services/order/dal/db"
	"gomall/services/order/dal/model"
	utils "gomall/services/order/utils/order"
	"strconv"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, _ error) {
	// TODO: Your code here...

	resp = new(order.PlaceOrderResp)
	resp.StatusCode = common.CodeServerBusy

	orders := make([]*model.Order, len(req.OrderItems))
	orderIds := make([]string, len(req.OrderItems))

	manager := utils.GetOrderIdGeneratorManager()
	var cost float32 = 0

	for i := 0; i < len(orders); i++ {
		oid, err := manager.GenerateId(req.UserId)
		if err != nil {
			zap.L().Error("generate order id fail", zap.Error(err))
			return
		}

		orders[i] = &model.Order{
			Oid:          oid,
			Uid:          req.UserId,
			UserCurrency: req.UserCurrency,

			Name:    req.Address.Name,
			Phone:   req.Address.Phone,
			Address: req.Address.Address,

			Pid:      req.OrderItems[i].Item.ProductId,
			Quantity: req.OrderItems[i].Item.Quantity,
			Cost:     req.OrderItems[i].Cost,

			Status: model.OrderStatusPending,
		}

		orderIds[i] = oid
		cost += req.OrderItems[i].Cost
	}

	err := db.AddOrders(ctx, orders)
	if err != nil {
		zap.L().Error("add order fail", zap.Error(err))
		return
	}

	oid, err := manager.GenerateId(req.UserId)
	if err != nil {
		zap.L().Error("generate order id fail", zap.Error(err))
		return
	}

	err = cache.SetPayId(ctx, orderIds, oid)
	if err != nil {
		zap.L().Error("set order id fail", zap.Error(err))
		return
	}

	resp.Order = &order.OrderResult{
		OrderId: oid,
		Cost:    cost,
	}
	resp.StatusCode = common.CodeSuccess

	return
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, _ error) {
	// TODO: Your code here...
	resp = new(order.ListOrderResp)
	resp.StatusCode = common.CodeServerBusy

	orders, err := db.GetOrders(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	resp.Orders = make([]*order.Order, len(orders))
	err = copier.Copy(&resp.Orders, orders)
	if err != nil {
		zap.L().Error("copy order fail", zap.Error(err))
		return nil, err
	}

	resp.StatusCode = common.CodeSuccess
	return
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, _ error) {
	// TODO: Your code here...
	resp = new(order.MarkOrderPaidResp)
	resp.StatusCode = common.CodeServerBusy

	orderIds := make([]string, 0)
	var err error
	orderIds, err = cache.GetPayId(ctx, req.OrderId)
	if err != nil {
		zap.L().Error("get order id fail", zap.Error(err))
		return
	}

	err = db.PutOrdersStatus(ctx, orderIds, model.OrderStatusPaid)
	if err != nil {
		zap.L().Error("put order status fail", zap.Error(err))
		return
	}

	resp.StatusCode = common.CodeSuccess

	return
}

func (s *OrderServiceImpl) MakeSureOrderExpired(ctx context.Context, req *order.MakeSureOrderExpiredReq) (resp *order.MakeSureOrderExpiredResp, _ error) {
	resp = new(order.MakeSureOrderExpiredResp)
	resp.StatusCode = common.CodeServerBusy
	resp.IsExpired = false

	_, err := cache.GetPayId(ctx, req.PayId)
	if err != nil {
		resp.IsExpired = true
	}

	return
}

func (s *OrderServiceImpl) MarkOrderShipped(ctx context.Context, req *order.MarkOrderShippedReq) (resp *order.MarkOrderShippedResp, _ error) {
	resp = new(order.MarkOrderShippedResp)
	resp.StatusCode = common.CodeServerBusy

	parseInt, err := strconv.ParseInt(req.OrderId, 10, 64)
	if err != nil {
		zap.L().Error("parse order id fail", zap.Error(err))
		return
	}

	order, err := db.GetOrderById(ctx, parseInt)
	if err != nil {
		zap.L().Error("get order id fail", zap.Error(err))
		return
	}
	if order.Status != model.OrderStatusPaid {
		resp.StatusCode = common.CodeOrderStatusErr
		return
	}

	err = db.PutOrderStatus(ctx, order.Oid, model.OrderStatusShipped)
	if err != nil {
		zap.L().Error("put order status fail", zap.Error(err))
		return
	}

	resp.StatusCode = common.CodeSuccess
	return
}

func (s *OrderServiceImpl) MarkOrderCompleted(ctx context.Context, req *order.MarkOrderCompletedReq) (resp *order.MarkOrderCompletedResp, _ error) {
	resp = new(order.MarkOrderCompletedResp)
	resp.StatusCode = common.CodeServerBusy

	parseInt, err := strconv.ParseInt(req.OrderId, 10, 64)
	if err != nil {
		zap.L().Error("parse order id fail", zap.Error(err))
		return
	}

	order, err := db.GetOrderById(ctx, parseInt)
	if err != nil {
		zap.L().Error("get order id fail", zap.Error(err))
		return
	}
	if order.Status != model.OrderStatusShipped {
		resp.StatusCode = common.CodeOrderStatusErr
		return
	}

	err = db.PutOrderStatus(ctx, order.Oid, model.OrderStatusCompleted)
	if err != nil {
		zap.L().Error("put order status fail", zap.Error(err))
		return
	}

	resp.StatusCode = common.CodeSuccess
	return
}

func NewOrderServiceImpl() *OrderServiceImpl {
	return &OrderServiceImpl{}
}
