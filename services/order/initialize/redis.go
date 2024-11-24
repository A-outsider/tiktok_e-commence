package initialize

import (
	"gomall/common/database"
	"gomall/services/order/config"
)

func initRedis() {
	svcContext.Client = database.NewRedisClientWithNoContext(config.GetConf().Redis)
}
