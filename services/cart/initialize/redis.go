package initialize

import (
	"gomall/common/database"
	"gomall/services/cart/config"
)

func initRedis() {
	svcContext.Client = database.NewRedisClientWithNoContext(config.GetConf().Redis)
}
