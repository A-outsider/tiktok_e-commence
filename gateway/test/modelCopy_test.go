package test

import (
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"gomall/gateway/types/req"
	rpcAuth "gomall/kitex_gen/auth"
	"testing"
)

func TestCopy(t *testing.T) {

	ctrlRequst := &req.LoginByPwdReq{
		Password:      "123456",
		Phone:         "12345678901",
		CaptchaID:     "12345678901",
		CaptchaAnswer: "123456",
	}

	expectedKitexReq := &rpcAuth.LoginByPwdReq{
		Password:      "123456",
		Phone:         "12345678901",
		CaptchaId:     "12345678901",
		CaptchaAnswer: "123456",
	}
	kitexReq := new(rpcAuth.LoginByPwdReq)

	// Act: 调用待测的 Copy 函数
	err := copier.Copy(kitexReq, ctrlRequst)

	// Assert: 验证复制结果是否符合预期
	assert.NoError(t, err, "复制过程中出现错误")
	assert.Equal(t, expectedKitexReq, kitexReq, "复制后的数据不匹配预期")
}
