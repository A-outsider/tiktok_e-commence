package initialize

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gomall/common/database"
	"gomall/services/user/config"
)

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

func InitRedis() {
	svcContext.RDB = new(RedisClient)
	svcContext.RDB.Client, svcContext.RDB.Ctx = database.NewRedisClient(config.GetConf().Redis)
}
