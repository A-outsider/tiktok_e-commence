package initialize

import (
	"gomall/common/database"
	"gomall/services/payment/config"
)

func initRedis() {
	svcContext.Client = database.NewRedisClientWithNoContext(config.GetConf().Redis)
}
