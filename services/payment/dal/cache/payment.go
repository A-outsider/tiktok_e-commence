package cache

import (
	"context"
	"gomall/common/utils/bloom_filter"
	"gomall/services/payment/initialize"
)

const (
	key     = "paymentId"
	bitsize = 100000
)

func AddBloom(ctx context.Context, value string) error {
	filter := bloom_filter.NewBloomFilter(initialize.GetRedis(), key, bitsize)
	return filter.Add(ctx, []byte(value))
}

func ExistBloom(ctx context.Context, value string) (bool, error) {
	filter := bloom_filter.NewBloomFilter(initialize.GetRedis(), key, bitsize)
	return filter.Exist(ctx, []byte(value))
}

func getPaymentId(value string) string {
	return "paymentId_" + value
}

func Set(ctx context.Context, value string) error {
	return initialize.GetRedis().Set(ctx, getPaymentId(value), 1, 0).Err()
}

func IsExist(ctx context.Context, value string) bool {
	return initialize.GetRedis().Exists(ctx, getPaymentId(value)).Err() == nil
}
