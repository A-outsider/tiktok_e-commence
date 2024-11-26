package cache

import (
	"context"
	"gomall/services/payment/initialize"
	"time"
)

func getPayKey(payId string) string {
	return "pay_" + payId
}

func SetPayId(ctx context.Context, orderId []string, payId string) error {
	err := initialize.GetRedis().SAdd(ctx, getPayKey(payId), orderId).Err()
	if err != nil {
		return err
	}

	return initialize.GetRedis().Expire(ctx, getPayKey(payId), time.Hour*24).Err()
}

func GetPayId(ctx context.Context, payId string) ([]string, error) {
	return initialize.GetRedis().SMembers(ctx, getPayKey(payId)).Result()
}
