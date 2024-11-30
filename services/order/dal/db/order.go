package db

import (
	"context"
	"gomall/services/order/dal/model"
	"gomall/services/order/initialize"
)

func AddOrders(ctx context.Context, orders []*model.Order) error {
	return initialize.GetMysql().WithContext(ctx).Create(orders).Error
}

func GetOrders(ctx context.Context, userId string) (orders []*model.Order, err error) {
	orders = make([]*model.Order, 0)
	err = initialize.GetMysql().WithContext(ctx).Model(&model.Order{}).Where("uid = ?", userId).Find(&orders).Error
	return
}

func GetOrdersBySeller(ctx context.Context, sellerId string) (orders []*model.Order, err error) {
	orders = make([]*model.Order, 0)
	pids := make([]string, 0)
	err = initialize.GetMysql().WithContext(ctx).Model(model.Product{}).Select("pid").Where("uid = ?", sellerId).Find(&pids).Error
	if err != nil {
		return
	}
	err = initialize.GetMysql().WithContext(ctx).Where("pid in ?", pids).Find(&orders).Error
	return
}

func GetOrderById(ctx context.Context, orderId int64) (order *model.Order, err error) {
	order = new(model.Order)
	err = initialize.GetMysql().WithContext(ctx).Model(&model.Order{}).Where("oid = ?", orderId).First(&order).Error
	return
}

func PutOrderStatus(ctx context.Context, orderId string, status int) error {
	return initialize.GetMysql().WithContext(ctx).Model(&model.Order{}).Where("oid = ?", orderId).Update("status", status).Error
}

func PutOrdersStatus(ctx context.Context, orderIds []string, status int) error {
	return initialize.GetMysql().WithContext(ctx).Model(&model.Order{}).Where("oid in (?)", orderIds).Update("status", status).Error
}
