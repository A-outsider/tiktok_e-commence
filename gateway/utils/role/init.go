package role

import (
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
	"gomall/gateway/config"
)

//-----------------------------------------------------用户权限管理--------------------------------------------------

// 定义角色关系
type Role int

// 用户角色
const (
	User = iota
	Seller
	Admin
)

var roleMap = map[Role]string{
	User:   "user",
	Seller: "seller",
	Admin:  "admin",
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
