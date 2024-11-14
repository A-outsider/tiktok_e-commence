// Package captcha 处理图片验证码逻辑
package captcha

import (
	"github.com/mojocn/base64Captcha"
	"gomall/common/utils/parse"
	"gomall/services/auth/config"
	"gomall/services/auth/dal/cache"
	"gomall/services/auth/initialize"
)

type PhCaptcha struct {
	Base64Captcha *base64Captcha.Captcha
}

// capt 内部使用的 Captcha 对象
var internalCaptcha *PhCaptcha

func NewCapt() *PhCaptcha {

	if internalCaptcha != nil {
		return internalCaptcha
	}

	// 初始化 Captcha 对象
	internalCaptcha = new(PhCaptcha)

	// 配置 base64Captcha 驱动信息
	driver := base64Captcha.NewDriverDigit(
		config.GetConf().PhotoCaptcha.Height,
		config.GetConf().PhotoCaptcha.Width,
		config.GetConf().PhotoCaptcha.Length,
		config.GetConf().PhotoCaptcha.MaxSkew,  // 倾斜角
		config.GetConf().PhotoCaptcha.DotCount, // 模糊点
	)

	// 实例化 base64Captcha 并赋值给内部使用的 capt 对象

	internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, internalCaptcha)

	return internalCaptcha
}

// GenerateCaptcha 生成图片验证码
func (c *PhCaptcha) GenerateCaptcha() (id string, b64s string, answer string, err error) {
	return c.Base64Captcha.Generate()
}

func (c *PhCaptcha) VerifyCaptcha(id string, answer string) (match bool) {
	return c.Base64Captcha.Verify(id, answer, false)
}

func (c *PhCaptcha) Set(id string, value string) error {
	_, err := initialize.GetRedis().SetWithTime(cache.GetPhotoCaptchaKey(id), value, parse.Duration(config.GetConf().PhotoCaptcha.Expire))
	return err
}

func (c *PhCaptcha) Get(id string, clear bool) string {
	val, err := initialize.GetRedis().Get(cache.GetPhotoCaptchaKey(id))
	if err != nil {
		return ""
	}

	if clear {
		initialize.GetRedis().Del(cache.GetPhotoCaptchaKey(id))
	}

	return val
}

func (c *PhCaptcha) Verify(id string, value string, clear bool) bool {
	val := c.Get(id, clear)
	return val == value
}
