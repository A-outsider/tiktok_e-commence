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
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisClient(val any) *RedisClient {
	r := new(Redis)

	err := mapstructure.Decode(val, r)
	if err != nil {
		log.Panic(fmt.Errorf("error decoding config to Redis struct: %v", err))
	}

	addr := fmt.Sprintf("%s:%d", r.Host, r.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: r.Password,
		DB:       0, // use default DB
	})

	ctx := context.Background()
	// Background返回一个非空的Context。它永远不会被取消，没有值，也没有截止日期。
	// 它通常由main函数、初始化和测试使用，并作为传入请求的顶级上下文
	// 设置一个5秒的超时
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Panic("redis connect ping failed, errs", zap.Error(err))
	}

	return &RedisClient{
		client: rdb,
		ctx:    ctx,
	}
}

func (r *RedisClient) SetWithTime(key string, value interface{}, dur time.Duration) (string, error) {
	return r.client.Set(r.ctx, key, value, dur).Result()
}

func (r *RedisClient) Set(key string, value interface{}) (string, error) {
	return r.client.Set(r.ctx, key, value, -1).Result()
}

func (r *RedisClient) Get(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("key cannot be empty")
	}
	return r.client.Get(r.ctx, key).Result()
}

func (r *RedisClient) Del(key string) {
	r.client.Del(r.ctx, key)
}

func (r *RedisClient) IncrWithTime(key string, times time.Duration) {
	r.client.Incr(r.ctx, key)
	r.client.Expire(r.ctx, key, times)
}

func (r *RedisClient) Keys(key string) []string {
	return r.client.Keys(r.ctx, key).Val()
}

func (r *RedisClient) TTL(key string) time.Duration {
	return r.client.TTL(r.ctx, key).Val()
}

func (r *RedisClient) ZRemRangeByRank(key string, start, stop int64) {
	r.client.ZRemRangeByRank(r.ctx, key, start, stop)
}

func (r *RedisClient) ZAdd(key string, score float64, member interface{}) {
	r.client.ZAdd(r.ctx, key, redis.Z{Score: score, Member: member})
}

func (r *RedisClient) ZCard(key string) int64 {
	return r.client.ZCard(r.ctx, key).Val()
}

func (r *RedisClient) ZScore(key string, member string) float64 {
	return r.client.ZScore(r.ctx, key, member).Val()
}

func (r *RedisClient) ZRem(key string, member ...interface{}) {
	r.client.ZRem(r.ctx, key, member...)
}
