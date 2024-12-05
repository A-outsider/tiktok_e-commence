package common

// 定义基础code状态码

const (
	// 成功
	CodeSuccess int64 = 1000
)

const (
	// 认证模块
	CodeInvalidParams int64 = 2001 + iota
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeNotLogin
	CodeInvalidCaptcha
	CodeInvalidCaptchaForm
	CodeInvalidEmailForm
	CodeInvalidPasswordForm
	CodeInvalidTokenForm
	CodeInvalidToken
	CodeInvalidRoleAdmin
	CodeInvalidDataUpdate
	CodeInvalidNewCaptcha
	CodeInvalidEmailWithUser
	CodeInvalidPhotoCaptcha
	CodeInvalidTokenExpired
	CodeUserALREADYLocked
	CodeVisitLimitExceeded
)

const (
	// 订单
	CodeOrderStatusErr int64 = 2201 + iota
)

const (
	// 支付
	CodePaySignatureVerifyFailed int64 = 2301 + iota
	CodePayMsgError
	CodePayIdExpired
	CodePayRepeat
)

const (
	// 其他错误  TODO 待规划
	CodeForbidden         int64 = 3001
	CodeServerBusy        int64 = 4001
	CodeRecordNotFound    int64 = 5001
	CodeRateLimitExceeded int64 = 6001
)

var Msg = map[int64]string{
	CodeSuccess: "success",

	// 认证模块
	CodeInvalidParams:        "请求参数错误",
	CodeUserExist:            "用户名已存在",
	CodeUserNotExist:         "用户不存在",
	CodeInvalidPassword:      "用户名或密码错误",
	CodeNotLogin:             "用户未登录",
	CodeInvalidCaptcha:       "手机号或验证码错误",
	CodeInvalidCaptchaForm:   "验证码格式错误",
	CodeInvalidEmailForm:     "用户邮箱格式错误",
	CodeInvalidPasswordForm:  "用户密码格式错误",
	CodeInvalidToken:         "无效的Token",
	CodeInvalidTokenForm:     "不合法的token格式",
	CodeInvalidRoleAdmin:     "用户权限不足",
	CodeInvalidDataUpdate:    "不合法的数据更新",
	CodeInvalidNewCaptcha:    "新邮箱的验证码错误",
	CodeInvalidEmailWithUser: "邮箱与用户信息不匹配",
	CodeInvalidPhotoCaptcha:  "图片验证码错误",
	CodeVisitLimitExceeded:   "访问流量达到限制",
	CodeInvalidTokenExpired:  "Token过期",
	CodeUserALREADYLocked:    "用户已被锁定",

	// 订单
	CodeOrderStatusErr: "订单当前状态错误",

	// 支付
	CodePaySignatureVerifyFailed: "支付签名验证失败",
	CodePayMsgError:              "支付失败",
	CodePayIdExpired:             "支付订单已过期",
	CodePayRepeat:                "重复支付",

	// 其他错误
	CodeForbidden:         "权限不足",
	CodeServerBusy:        "服务繁忙",
	CodeRecordNotFound:    "未查询到该记录",
	CodeRateLimitExceeded: "操作频率过快 ,请稍后再试",
}
