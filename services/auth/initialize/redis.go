package initialize

import (
	"gomall/common/database"
	"gomall/services/auth/config"
)

func InitRedis() {
	svcContext.RDB = database.NewRedisClient(config.GetConf().Redis)
}
