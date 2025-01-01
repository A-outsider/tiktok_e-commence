package initialize

import (
	"gomall/common/database"
	"gomall/gateway/config"
)

func initRedis() {
	svcContext.Client = database.NewRedisClientWithNoContext(config.GetConf().Redis)
}
