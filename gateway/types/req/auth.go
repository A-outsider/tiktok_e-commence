package req

// 登录
type LoginByCodeReq struct {
	Phone    string `form:"phone" binding:"required,phone"`
	AuthCode string `form:"auth_code"`
}

type LoginByCodeResp struct {
	Token        string `form:"token"`
	RefreshToken string `form:"refresh_token"`
}

type LoginByPwdReq struct {
	Phone         string `form:"phone" `
	Password      string `form:"password" `
	CaptchaID     string `form:"captcha_id,omitempty"`
	CaptchaAnswer string `form:"captcha_answer,omitempty"`
}

type LoginByPwdResp struct {
	Token        string `form:"token"`
	RefreshToken string `form:"refresh_token"`
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

type RefreshTokenResp struct {
	Token        string `form:"token"`
	RefreshToken string `form:"refresh_token"`
}

// 发送验证码
type SendPhoneCodeReq struct {
	Phone string `form:"phone" binding:"required,phone"`
}

type SendEmailCodeReq struct {
	Email string `form:"email" binding:"required,email"`
}

// 图片
type ShowPhotoCaptchaResp struct {
	CaptchaId  string `form:"captcha_id"`
	CaptchaImg string `form:"captcha_img"`
}
