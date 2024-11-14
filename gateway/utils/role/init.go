package role

import (
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
	"gomall/gateway/config"
)

//-----------------------------------------------------用户权限管理--------------------------------------------------

// 定义角色关系
type Role int

const (
	// 未认证
	MORTAL Role = 0
	// 普通买家
	USER Role = 1
	// 普通卖家
	Seller Role = 2
	// 审核
	DEVELOP Role = 6
	// 管理员
	ADMIN Role = 9
	// 超级管理员
	ROOT Role = 10
)

var roleMap = map[Role]string{
	MORTAL:  "mortal",
	USER:    "user",
	Seller:  "seller",
	DEVELOP: "develop",
	ADMIN:   "admin",
	ROOT:    "root",
}

var enforcer *casbin.Enforcer

func InitCasbin() {
	var err error
	enforcer, err = casbin.NewEnforcer(config.GetConf().Role.Model, config.GetConf().Role.Policy)
	if err != nil {
		zap.L().Error("初始化casbin错误，errs：" + err.Error())
	}
}

// 用户身份int转对应的string
func GetRoleString(r int64) string {
	return roleMap[Role(r)]
}
