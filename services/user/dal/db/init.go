package db

import (
	"errors"
	"gomall/common/database"
	"gomall/services/user/dal/model"
	"gomall/services/user/utils/password"
	"gorm.io/gorm"
)

// 数据库初始化

func InitDb(config any) error {

	err := database.InitMySQL(config)
	if err != nil {
		return err
	}

	// 关联表
	err = database.GetDB().AutoMigrate(&model.User{})
	if err != nil {
		return err
	}

	// 创建超级管理员
	_, err = GetUser(1)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = InsertUser(&model.User{
			ID:       "1",
			Phone:    "12345678901",
			Username: "tiktok_admin",
			Password: password.Encrypt("lijialang666"),
			Role:     model.ConstRoleOfAdmin,
		})
	}

	return err

}
