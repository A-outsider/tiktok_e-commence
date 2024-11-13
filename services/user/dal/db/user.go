package db

import (
	"gomall/common/database"
	"gomall/services/user/dal/model"
)

func GetUser(id int64) (*model.User, error) {
	user := new(model.User)
	err := database.GetDB().Where("id = ?", id).First(user).Error
	return user, err
}

// InsertUser 插入用户信息
func InsertUser(user *model.User) error {
	return database.GetDB().Create(user).Error
}
