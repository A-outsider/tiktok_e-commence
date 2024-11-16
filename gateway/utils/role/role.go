package role

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"go.uber.org/zap"
	"gomall/gateway/rpc"
	"gomall/gateway/types/resp"
	auth "gomall/kitex_gen/auth"
	"sync"
)

var checkLock sync.Mutex

// 验证权限
func CheckAdmin(c context.Context, ctx *app.RequestContext, userId string) (StatusCode int64) {

	result, _ := rpc.GetAuthClient().GetUserAdmin(c, &auth.CheckAdminReq{UserId: userId})
	if result == nil || result.StatusCode == 0 {
		return resp.CodeServerBusy
	}
	if result.StatusCode != resp.CodeSuccess {
		return result.StatusCode
	}

	AdminRole := GetRoleString(result.Role)

	return check(userId, AdminRole, ctx.FullPath(), string(ctx.Request.Method()))
}

// 目前策略模型 , 管理员 继承 卖家接口 , 卖家 继承 普通用户接口
func check(userId string, sub, obj, act string) (StatusCode int64) {
	checkLock.Lock()
	defer checkLock.Unlock()

	ok, _ := enforcer.Enforce(sub, obj, act) // sub主体 , obj对象 , act动作
	if ok {
		return resp.CodeSuccess
	}
	zap.L().Error(fmt.Sprintf("权限不足,用户ID：%d，角色：%s，路径：%s，请求方法：%s", userId, sub, obj, act))
	return resp.CodeInvalidRoleAdmin
}
