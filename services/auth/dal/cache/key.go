package cache

import "fmt"

// Key 前缀常量
const (
	EmailCaptchaPrefix     = "captcha:"
	SendCaptchaIntervalKey = "send_captcha_interval:"
	CaptchaPhotoPrefix     = "photoCaptcha:"
	UserRolePrefix         = "user_role:"
	RefreshTokenKey        = "refresh_token_key:"
	ErrorPsdLimitKey       = "error_psd_limit:"
	PhoneCaptchaPrefix     = "phone_captcha:"
)

// 邮箱验证码
func GetEmailKey(email string) string {
	return fmt.Sprintf(EmailCaptchaPrefix+"%s", email)
}

// 发送验证码时间间隔
func GetSendCaptchaIntervalKey(email string) string {
	return fmt.Sprintf(SendCaptchaIntervalKey+"%s", email)
}

// 图片验证码
func GetPhotoCaptchaKey(email string) string {
	return fmt.Sprintf(CaptchaPhotoPrefix+"%s", email)
}

// 用户角色
func GetUserRoleKey(userid string) string {
	return fmt.Sprintf(UserRolePrefix+"%s", userid)
}

// 刷新token缓存标识符
func GetRefreshTokenKey(userid string) string {
	return fmt.Sprintf(RefreshTokenKey+"%s", userid)
}

// 错误密码尝试次数
func GetErrorPsdLimitKey(id string) string {
	return fmt.Sprintf(ErrorPsdLimitKey+"%s", id)
}

// 手机验证码
func GetPhoneCodeKey(phone string) string {
	return fmt.Sprintf(PhoneCaptchaPrefix+"%s", phone)
}
