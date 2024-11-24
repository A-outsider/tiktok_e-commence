package handler

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gomall/gateway/types/resp/common"
	order "gomall/kitex_gen/order"
	"gomall/services/order/dal/db"
	"gomall/services/order/dal/model"
	utils "gomall/services/order/utils/order"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// TODO: Your code here...

	resp = new(order.PlaceOrderResp)
	resp.StatusCode = common.CodeServerBusy

	orders := make([]*model.Order, len(req.OrderItems))
	orderResults := make([]*order.OrderResult, len(req.OrderItems))

	manager := utils.GetOrderIdGeneratorManager()

	for i := 0; i < len(orders); i++ {
		var oid string
		oid, err = manager.GenerateId(req.UserId)
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

		orderResults[i] = &order.OrderResult{
			OrderId: orders[i].Oid,
		}
	}

	err = db.AddOrders(ctx, orders)
	if err != nil {
		zap.L().Error("add order fail", zap.Error(err))
		return
	}

	resp.Order = orderResults
	resp.StatusCode = common.CodeSuccess

	return
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
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
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// TODO: Your code here...
	resp = new(order.MarkOrderPaidResp)
	resp.StatusCode = common.CodeServerBusy

	err = db.PutOrderStatus(ctx, req.OrderId, model.OrderStatusPaid)
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
