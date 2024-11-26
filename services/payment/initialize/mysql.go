package initialize

import (
	"gomall/common/database"
	"gomall/services/payment/config"
	"gomall/services/payment/dal/model"
)

func initMysql() {
	svcContext.DB = database.NewMySQL(config.GetConf().Mysql) // 关联表
	err := svcContext.DB.AutoMigrate(&model.Order{})
	if err != nil {
		panic(err)
	}
}
