package req

// 登录
type LoginByCodeReq struct {
	Phone    string `form:"phone" binding:"required,phone"`
	AuthCode string `form:"auth_code"`
}

type LoginByPwdReq struct {
	Phone         string `form:"phone" `
	Password      string `form:"password" `
	CaptchaID     string `form:"captcha_id,omitempty"`
	CaptchaAnswer string `form:"captcha_answer,omitempty"`
}

// 注册
type RegisterReq struct {
	Phone    string `form:"phone" binding:"required,phone"`
	AuthCode string `form:"auth_code" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// 刷新token
type RefreshTokenReq struct {
	RefreshToken string `form:"refresh_token" binding:"required"`
}

// 发送验证码
type SendPhoneCodeReq struct {
	Phone string `form:"phone" binding:"required,phone"`
}

type SendEmailCodeReq struct {
	Email string `form:"email" binding:"required,email"`
}

// 修改为卖家
type ModifyUserToSellerReq struct {
	UserId string `form:"user_id" binding:"required"`
}
