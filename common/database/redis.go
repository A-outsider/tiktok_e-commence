package database

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var (
	redisClient = new(redis.Client)
	ctx         = context.Background()
)

// 注册redis
func Init() {
	conf := config.Get().Redis
	addr := fmt.Sprintf(conf.Host + ":" + strconv.Itoa(conf.Port))
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Password,
		DB:       0, // use default DB
	})

	// Background返回一个非空的Context。它永远不会被取消，没有值，也没有截止日期。
	// 它通常由main函数、初始化和测试使用，并作为传入请求的顶级上下文
	// 设置一个5秒的超时

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		zap.L().Error("redis connect ping failed, errs", zap.Error(err))
		redisClient = nil
	}

}

func SetWithTime(key string, value interface{}, dur time.Duration) (string, error) {
	return redisClient.Set(ctx, key, value, dur).Result()
}

func Set(key string, value interface{}) (string, error) {
	return redisClient.Set(ctx, key, value, -1).Result()
}

func Get(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("key cannot be empty")
	}
	return redisClient.Get(ctx, key).Result()
}

func Del(key string) {
	redisClient.Del(ctx, key)
}

func IncrWithTime(key string, times time.Duration) {
	redisClient.Incr(ctx, key)
	redisClient.Expire(ctx, key, times)
}

func Keys(key string) []string {
	return redisClient.Keys(ctx, key).Val()
}

func TTL(key string) time.Duration {
	return redisClient.TTL(ctx, key).Val()
}

// 移除有序集中指定排名区间内的所有成员
func ZRemRangeByRank(key string, start, stop int64) {
	redisClient.ZRemRangeByRank(ctx, key, start, stop)
}

// 向有序集合插入数据
func ZAdd(key string, score float64, member interface{}) {
	redisClient.ZAdd(ctx, key, redis.Z{Score: score, Member: member})
}

// 有序集合中的数量
func ZCard(key string) int64 {
	return redisClient.ZCard(ctx, key).Val()
}

// 有序集中成员的分数值
func ZScore(key string, member string) float64 {
	return redisClient.ZScore(ctx, key, member).Val()
}

// 向有序集合插入数据
func ZRem(key string, member ...interface{}) {
	redisClient.ZRem(ctx, key, member...)
}
