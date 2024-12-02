package initialize

import (
	"gomall/common/database"
	"gomall/services/product/config"
)

func initRedis() {
	svcContext.Client = database.NewRedisClientWithNoContext(config.GetConf().Redis)
}
