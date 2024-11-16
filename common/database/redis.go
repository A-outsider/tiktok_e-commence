package database

import (
	"context"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gomall/services/auth/initialize"
	"log"
)

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

func NewRedisClient(val any) *initialize.RedisClient {
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

	return &initialize.RedisClient{
		Client: rdb,
		Ctx:    ctx,
	}
}
