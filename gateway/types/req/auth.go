package req

// 登录
type LoginByCodeReq struct {
	Phone    string `json:"phone" binding:"required,phone"`
	AuthCode string `json:"auth_code" binding:"required"`
}

type LoginByCodeResp struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginByPwdReq struct {
	Phone         string `json:"phone" binding:"required,phone"`
	Password      string `json:"password" binding:"required"`
	CaptchaID     string `json:"captcha_id,omitempty"`
	CaptchaAnswer string `json:"captcha_answer,omitempty"`
}

type LoginByPwdResp struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

// 注册
type RegisterReq struct {
	Phone    string `json:"phone" binding:"required,phone"`
	AuthCode string `json:"auth_code" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 刷新token
type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshTokenResp struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

// 发送验证码
type SendPhoneCodeReq struct {
	Phone string `json:"phone" binding:"required,phone"`
}

type SendEmailCodeReq struct {
	Email string `json:"email" binding:"required,email"`
}

// 图片
type ShowPhotoCaptchaResp struct {
	CaptchaId  string `json:"captcha_id"`
	CaptchaImg string `json:"captcha_img"`
}
