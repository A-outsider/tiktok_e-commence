package resp

type LoginByCodeResp struct {
	Token        string `form:"token"`
	RefreshToken string `form:"refresh_token"`
}

type LoginByPwdResp struct {
	Token        string `form:"token"`
	RefreshToken string `form:"refresh_token"`
}

type RefreshTokenResp struct {
	Token        string `form:"token"`
	RefreshToken string `form:"refresh_token"`
}

// 图片
type ShowPhotoCaptchaResp struct {
	CaptchaId  string `form:"captcha_id"`
	CaptchaImg string `form:"captcha_img"`
}
