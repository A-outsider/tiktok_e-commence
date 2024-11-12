package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

var (
	rdb = new(redis.Client)
	ctx = context.Background()
)

func GetRdbCli() *redis.Client {
	return rdb
}

// 注册redis
func InitRedis(val any) {
	r := new(Redis)

	err := mapstructure.Decode(val, r)
	if err != nil {
		log.Panic(fmt.Errorf("error decoding config to Redis struct: %v", err))
	}

	addr := fmt.Sprintf(r.Host + ":" + r.Port)
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: r.Password,
		DB:       0, // use default DB
	})

	// Background返回一个非空的Context。它永远不会被取消，没有值，也没有截止日期。
	// 它通常由main函数、初始化和测试使用，并作为传入请求的顶级上下文
	// 设置一个5秒的超时

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Panic("redis connect ping failed, errs", zap.Error(err))
	}

	// TODO : 为redis启动链路追踪
}

func SetWithTime(key string, value interface{}, dur time.Duration) (string, error) {
	return rdb.Set(ctx, key, value, dur).Result()
}

func Set(key string, value interface{}) (string, error) {
	return rdb.Set(ctx, key, value, -1).Result()
}

func Get(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("key cannot be empty")
	}
	return rdb.Get(ctx, key).Result()
}

func Del(key string) {
	rdb.Del(ctx, key)
}

func IncrWithTime(key string, times time.Duration) {
	rdb.Incr(ctx, key)
	rdb.Expire(ctx, key, times)
}

func Keys(key string) []string {
	return rdb.Keys(ctx, key).Val()
}

func TTL(key string) time.Duration {
	return rdb.TTL(ctx, key).Val()
}

// 移除有序集中指定排名区间内的所有成员
func ZRemRangeByRank(key string, start, stop int64) {
	rdb.ZRemRangeByRank(ctx, key, start, stop)
}

// 向有序集合插入数据
func ZAdd(key string, score float64, member interface{}) {
	rdb.ZAdd(ctx, key, redis.Z{Score: score, Member: member})
}

// 有序集合中的数量
func ZCard(key string) int64 {
	return rdb.ZCard(ctx, key).Val()
}

// 有序集中成员的分数值
func ZScore(key string, member string) float64 {
	return rdb.ZScore(ctx, key, member).Val()
}

// 向有序集合插入数据
func ZRem(key string, member ...interface{}) {
	rdb.ZRem(ctx, key, member...)
}
