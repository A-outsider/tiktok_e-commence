package db

import (
	"gomall/services/auth/dal/model"
	"gomall/services/auth/initialize"
)

func GetUser(id string) (*model.User, error) {
	user := new(model.User)
	err := initialize.GetMysql().Where("id = ?", id).First(user).Error
	return user, err
}

// InsertUser 插入用户信息
func InsertUser(user *model.User) error {
	return initialize.GetMysql().Create(user).Error
}

func SelectUserByPhone(phone string) (user *model.User, err error) {
	user = new(model.User)
	err = initialize.GetMysql().Where("phone = ?", phone).First(user).Error
	return
}

func GetUserByID(userId string) (user *model.User, err error) {
	user = new(model.User)
	return user, initialize.GetMysql().Where("id = ?", userId).First(user).Error
}

func UpdateUserRoleTOSeller(userId string) error {

	return initialize.GetMysql().Model(&model.User{}).Where("id = ?", userId).Update("role", 1).Error // 1 为卖家
}
