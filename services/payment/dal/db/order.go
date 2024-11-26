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

func PutOrderStatus(ctx context.Context, orderId string, status int) error {
	return initialize.GetMysql().WithContext(ctx).Model(&model.Order{}).Where("order_id = ?", orderId).Update("status", status).Error
}

func PutOrdersStatus(ctx context.Context, orderIds []string, status int) error {
	return initialize.GetMysql().WithContext(ctx).Model(&model.Order{}).Where("order_id in (?)", orderIds).Update("status", status).Error
}
