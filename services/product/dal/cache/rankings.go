package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gomall/services/product/dal/model"
	"gomall/services/product/initialize"
	"strconv"
)

func getRankingsKey() string {
	return "hot_product_rankings"
}

func getProductKey(pid string) string {
	return "product_info:" + pid
}

func AddProductToRanking(ctx context.Context, product *model.Product) error {
	rdb := initialize.GetRedis()
	err := rdb.ZAdd(ctx, getRankingsKey(), redis.Z{Score: 0.0, Member: product.Pid}).Err()
	if err != nil {
		return err
	}

	return setProductDetails(rdb, ctx, product.Pid, product.Name, product.Picture, product.Price)
}

func IncrProductHotness(ctx context.Context, pid string) error {
	return initialize.GetRedis().ZIncrBy(ctx, getRankingsKey(), 1, pid).Err()
}

func GetRankingsWithDetails(ctx context.Context, topN int) ([]map[string]string, error) {
	// 获取前N名商品的ID
	rdb := initialize.GetRedis()
	productIDs, err := rdb.ZRevRangeWithScores(context.Background(), getRankingsKey(), 0, int64(topN-1)).Result()
	if err != nil {
		return nil, err
	}

	// 获取每个商品的详细信息
	var productDetailsList []map[string]string
	for _, productID := range productIDs {
		details, err := getProductDetails(ctx, rdb, productID.Member.(string))
		if err != nil {
			return nil, err
		}

		details["pid"] = productID.Member.(string)
		details["score"] = strconv.FormatFloat(productID.Score, 'e', 2, 64)
		productDetailsList = append(productDetailsList, details)
	}

	return productDetailsList, nil
}

func setProductDetails(rdb *redis.Client, ctx context.Context, productID, productName, imageURL string, price float32) error {
	// 使用HSET命令将商品的详细信息存储到哈希表中
	_, err := rdb.HSet(ctx, getProductKey(productID), map[string]interface{}{
		"name":    productName,
		"picture": imageURL,
		"price":   price,
	}).Result()
	return err
}

func getProductDetails(ctx context.Context, rdb *redis.Client, productID string) (map[string]string, error) {
	// 使用HGETALL命令获取商品的详细信息
	details, err := rdb.HGetAll(ctx, getProductKey(productID)).Result()
	if err != nil {
		return nil, err
	}
	return details, nil
}
