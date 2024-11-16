package initialize

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gomall/common/database"
	"gomall/services/auth/config"
)

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

func InitRedis() {
	svcContext.RDB = database.NewRedisClient(config.GetConf().Redis)
}
