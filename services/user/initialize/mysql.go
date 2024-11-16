package initialize

import (
	"gomall/common/database"
	"gomall/services/user/config"
	"gomall/services/user/dal/model"
)

func initMysql() {
	svcContext.DB = database.NewMySQL(config.GetConf().Mysql) // 关联表
	err := svcContext.DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
