package test

import (
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	user "gomall/kitex_gen/user"
	"gomall/services/user/dal/model"
	"testing"
)

// TestCopyUser 测试复制用户信息的函数
func TestCopyUser(t *testing.T) {
	// 初始化mock对象

	mockReq := &user.ModifyUserInfoReq{
		Id:       "123",
		Name:     "test",
		Gender:   1,
		Birthday: "1990-01-01",
	}

	//_ := model.User{
	//	ID:       "123",
	//	Name:     "test",
	//	Gender:   1,
	//	Birthday: "1990-01-01",
	//}

	u := new(model.User)
	// 开始测试
	err := copier.Copy(u, mockReq) // 假设CopyUser是待测函数

	// 验证没有错误
	assert.NoError(t, err)
	assert.NotNil(t, u)
	//assert.Equal(t, u, expectu, "复制后的数据不匹配预期")
}
