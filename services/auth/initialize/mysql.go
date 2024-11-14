package initialize

import (
	"gomall/common/database"
	"gomall/services/auth/config"
	"gomall/services/auth/dal/model"
)

func initMysql() {
	svcContext.DB = database.NewMySQL(config.GetConf().Mysql) // 关联表
	err := svcContext.DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
