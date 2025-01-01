package initialize

import (
	"gomall/common/database"
	"gomall/services/auth/config"
)

func initRedisWithNoContext() {
	svcContext.Client = database.NewRedisClientWithNoContext(config.GetConf().Redis)
}
