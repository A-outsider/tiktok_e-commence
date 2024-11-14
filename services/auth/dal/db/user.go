package db

import (
	"gomall/common/database"
	"gomall/services/auth/dal/model"
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

func SelectUserByPhone(phone string) (user *model.User, err error) {
	user = new(model.User)
	err = database.GetDB().Where("phone = ?", phone).First(user).Error
	return
}

func GetUserByID(userId string) (user *model.User, err error) {
	user = new(model.User)
	return user, database.GetDB().Where("id = ?", userId).First(user).Error
}
