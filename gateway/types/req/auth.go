package req

type LoginReq struct {
	Phone         string `json:"phone" binding:"required,phone"`
	Password      string `json:"password" binding:"required,"`
	CaptchaID     string `json:"captcha_id,omitempty"`
	CaptchaAnswer string `json:"captcha_answer,omitempty"`
}

type LoginResp struct {
}
