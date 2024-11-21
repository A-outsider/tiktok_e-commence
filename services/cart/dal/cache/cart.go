package cache

import (
	"context"
	cart "gomall/kitex_gen/cart"
	"gomall/services/cart/initialize"
)

func getCartKey(uid string) string {
	return "cart_" + uid
}

func AddItem(ctx context.Context, uid string, pid string, quantity int64) error {
	return initialize.GetRedis().ZIncrBy(ctx, getCartKey(uid), float64(quantity), pid).Err()
}

func GetItems(ctx context.Context, uid string) ([]*cart.CartItem, error) {
	results, err := initialize.GetRedis().ZRangeWithScores(ctx, getCartKey(uid), 0, -1).Result()
	if err != nil {
		return nil, err
	}

	res := make([]*cart.CartItem, len(results))
	for i := 0; i < len(results); i++ {
		res[i] = &cart.CartItem{
			ProductId: results[i].Member.(string),
			Quantity:  int64(results[i].Score),
		}
	}

	return res, nil
}

func DeleteItems(ctx context.Context, uid string) error {
	return initialize.GetRedis().Del(ctx, getCartKey(uid)).Err()
}
