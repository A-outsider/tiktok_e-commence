package cache

import (
	"errors"
	"github.com/redis/go-redis/v9"
	"gomall/services/user/initialize"
	"time"
)

func SetWithTime(key string, value interface{}, dur time.Duration) (string, error) {
	return initialize.GetRedis().Client.Set(initialize.GetRedis().Ctx, key, value, dur).Result()
}

func Set(key string, value interface{}) (string, error) {
	return initialize.GetRedis().Client.Set(initialize.GetRedis().Ctx, key, value, -1).Result()
}

func Get(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("key cannot be empty")
	}
	return initialize.GetRedis().Client.Get(initialize.GetRedis().Ctx, key).Result()
}

func Del(key string) {
	initialize.GetRedis().Client.Del(initialize.GetRedis().Ctx, key)
}

func IncrWithTime(key string, times time.Duration) {
	initialize.GetRedis().Client.Incr(initialize.GetRedis().Ctx, key)
	initialize.GetRedis().Client.Expire(initialize.GetRedis().Ctx, key, times)
}

func Keys(key string) []string {
	return initialize.GetRedis().Client.Keys(initialize.GetRedis().Ctx, key).Val()
}

func TTL(key string) time.Duration {
	return initialize.GetRedis().Client.TTL(initialize.GetRedis().Ctx, key).Val()
}

func ZRemRangeByRank(key string, start, stop int64) {
	initialize.GetRedis().Client.ZRemRangeByRank(initialize.GetRedis().Ctx, key, start, stop)
}

func ZAdd(key string, score float64, member interface{}) {
	initialize.GetRedis().Client.ZAdd(initialize.GetRedis().Ctx, key, redis.Z{Score: score, Member: member})
}

func ZCard(key string) int64 {
	return initialize.GetRedis().Client.ZCard(initialize.GetRedis().Ctx, key).Val()
}

func ZScore(key string, member string) float64 {
	return initialize.GetRedis().Client.ZScore(initialize.GetRedis().Ctx, key, member).Val()
}

func ZRem(key string, member ...interface{}) {
	initialize.GetRedis().Client.ZRem(initialize.GetRedis().Ctx, key, member...)
}
