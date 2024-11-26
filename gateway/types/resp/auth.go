package resp

type LoginByCodeResp struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginByPwdResp struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResp struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

// 图片
type ShowPhotoCaptchaResp struct {
	CaptchaId  string `json:"captcha_id"`
	CaptchaImg string `json:"captcha_img"`
}
