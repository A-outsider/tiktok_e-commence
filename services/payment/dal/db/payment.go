package db

import (
	"context"
	"gomall/services/payment/dal/model"
	"gomall/services/payment/initialize"
)

func CreatePaymentId(ctx context.Context, payId string) error {
	payment := model.Payment{
		PayId: payId,
	}
	return initialize.GetMysql().WithContext(ctx).Create(&payment).Error
}

func IsHavePayment(ctx context.Context, payId string) bool {
	payment := model.Payment{}
	return initialize.GetMysql().WithContext(ctx).First(&payment, "pay_id = ?", payId).Error == nil
}
