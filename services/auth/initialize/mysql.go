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

//// 创建超级管理员
//_, err = db.GetUser(1)
//if errors.Is(err, gorm.ErrRecordNotFound) {
//err = db.InsertUser(&model.User{
//ID:       "1",
//Phone:    "12345678901",
//Username: "tiktok_admin",
//Password: password.Encrypt("lijialang666"),
//Role:     model.ConstRoleOfAdmin,
//})
//}
//
