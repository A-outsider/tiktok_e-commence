package initialize

import (
	"gomall/common/database"
	"gomall/services/auth/config"
)

func initSms() {
	svcContext.SMS = database.NewSms(config.GetConf().Phone)
}
