package db

import (
	"gomall/services/user/dal/model"
	"gomall/services/user/initialize"
)

func GetUser(id string) (*model.User, error) {
	user := new(model.User)
	err := initialize.GetMysql().Where("id = ?", id).First(user).Error
	return user, err
}

func GetUserByID(userId string) (user *model.User, err error) {
	user = new(model.User)
	return user, initialize.GetMysql().Where("id = ?", userId).First(user).Error
}

func ModifyUserInfo(Id string, user *model.User) error { // TODO: 考虑细化更新响应的参数
	return initialize.GetMysql().Model(&model.User{}).Where("id = ?", Id).Updates(user).Error
}

func DeleteUserById(Id string) error {
	return initialize.GetMysql().Where("id = ?", Id).Delete(&model.User{}).Error
}

// address
func GetAddressList(Id string) ([]*model.Address, error) {
	var addressList []*model.Address
	err := initialize.GetMysql().Where("uid = ?", Id).Find(&addressList).Error
	return addressList, err
}

func GetAddressById(id, aid string) (*model.Address, error) {
	address := new(model.Address)
	return address, initialize.GetMysql().Where("uid = ? and aid = ?", id, aid).First(address).Error
}

func CreateAddress(address *model.Address) error {
	return initialize.GetMysql().Create(address).Error
}

func ModifyAddress(id, aid string, address *model.Address) error { // TODO: 考虑细化更新响应的参数
	return initialize.GetMysql().Model(&model.Address{}).Where("uid = ? and aid = ?", id, aid).Updates(address).Error
}

func DeleteAddress(id, aid string) bool {
	result := initialize.GetMysql().Where("uid = ? and aid = ?", id, aid).Delete(&model.Address{})
	if result.RowsAffected == 0 {
		return false
	}
	return true
}
