package handler

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"gomall/common/utils/parse"
	"gomall/common/utils/random"
	"gomall/gateway/types/resp"
	auth "gomall/kitex_gen/auth"
	"gomall/services/auth/config"
	"gomall/services/auth/dal/cache"
	"gomall/services/auth/dal/db"
	"gomall/services/auth/dal/model"
	"gomall/services/auth/utils/captcha"
	"gomall/services/auth/utils/mail"
	"gomall/services/auth/utils/password"
	"gomall/services/auth/utils/sms"
	"gomall/services/auth/utils/token"
	"gorm.io/gorm"
	"strconv"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// NewAuthServiceImpl creates a new instance of AuthServiceImpl.
func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{}
}

// LoginByCode implements the AuthService interface.
func (s *AuthServiceImpl) LoginByCode(ctx context.Context, req *auth.LoginByCodeReq) (res *auth.LoginByCodeResp, err error) {
	res = new(auth.LoginByCodeResp)
	res.StatusCode = resp.CodeServerBusy

	// 校验用户是否存在
	user := new(model.User)
	user, err = db.SelectUserByPhone(req.GetPhone())

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || len(user.ID) == 0 {
			res.StatusCode = resp.CodeUserNotExist
			return
		} else {
			zap.L().Error("CheckPassword fail", zap.Error(err))
			return
		}
	}

	// 校验验证码
	code, err := cache.Get(cache.GetPhoneCodeKey(req.GetPhone()))
	if err != nil {
		return nil, err
	}

	if code != req.GetCode() {
		res.StatusCode = resp.CodeInvalidCaptcha
		return
	}

	// 删缓存
	go cache.Del(cache.GetRefreshTokenKey(user.ID))

	// 生成验证token
	if res.Token, err = token.GenerateAccessToken(user.ID); err != nil || len(res.Token) == 0 {
		zap.L().Error("验证token生成失败", zap.Error(err))
		return
	}

	// 生成刷新token
	if res.RefreshToken, err = token.GenerateRefreshToken(user.ID); err != nil || len(res.RefreshToken) == 0 {
		zap.L().Error("刷新token生成失败", zap.Error(err))
		return
	}

	// 存入缓存
	_, err = cache.SetWithTime(cache.GetRefreshTokenKey(user.ID), res.RefreshToken, parse.Duration(config.GetConf().Jwt.RefreshExpireTime))
	if err != nil {
		zap.L().Error("redis.Set fail", zap.Error(err))
		return
	}

	// success
	res.StatusCode = resp.CodeSuccess

	return
}

// LoginByPwd implements the AuthService interface.
func (s *AuthServiceImpl) LoginByPwd(ctx context.Context, req *auth.LoginByPwdReq) (res *auth.LoginByPwdResp, err error) {
	res = new(auth.LoginByPwdResp)
	res.StatusCode = resp.CodeServerBusy

	// 校验用户是否存在
	user := new(model.User)
	user, err = db.SelectUserByPhone(req.GetPhone())

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || len(user.ID) == 0 {
			res.StatusCode = resp.CodeUserNotExist
			return
		} else {
			zap.L().Error("CheckPassword fail", zap.Error(err))
			return
		}
	}

	// 校验图片验证码
	if ok := captcha.NewCapt().VerifyCaptcha(req.GetCaptchaId(), req.GetCaptchaAnswer()); !ok {
		res.StatusCode = resp.CodeInvalidCaptcha
		return
	}

	// 密码错误次数限制
	failures, _ := cache.Get(cache.GetErrorPsdLimitKey(user.ID))

	var f int
	if f, err = strconv.Atoi(failures); err == nil && f >= config.GetConf().Password.ErrorLimit {
		res.StatusCode = resp.CodeUserALREADYLocked
		return
	}

	// 密码校验
	if password.Encrypt(req.GetPassword()) != user.Password {
		cache.IncrWithTime(cache.GetErrorPsdLimitKey(user.ID), parse.Duration(config.GetConf().Password.ErrorLockTime))
		res.StatusCode = resp.CodeInvalidPassword
		return
	}

	// 登录成功 , 删除缓存
	cache.Del(cache.GetErrorPsdLimitKey(user.ID))
	cache.Del(cache.GetRefreshTokenKey(user.ID))

	// 生成验证token
	if res.Token, err = token.GenerateAccessToken(user.ID); err != nil || len(res.Token) == 0 {
		zap.L().Error("验证token生成失败", zap.Error(err))
		return
	}
	// 生成刷新token
	if res.RefreshToken, err = token.GenerateRefreshToken(user.ID); err != nil || len(res.RefreshToken) == 0 {
		zap.L().Error("刷新token生成失败", zap.Error(err))
		return
	}

	// 存入缓存
	_, err = cache.SetWithTime(cache.GetRefreshTokenKey(user.ID), res.RefreshToken, parse.Duration(config.GetConf().Jwt.RefreshExpireTime))
	if err != nil {
		zap.L().Error("redis.Set fail", zap.Error(err))
		return
	}

	res.StatusCode = resp.CodeSuccess
	return
}

// Register implements the AuthService interface.
const defaultAvatar = "./data/" // TODO : 写入配置文件
func (s *AuthServiceImpl) Register(ctx context.Context, req *auth.RegisterReq) (res *auth.RegisterResp, err error) {
	res = new(auth.RegisterResp)
	res.StatusCode = resp.CodeServerBusy

	// 校验手机验证码	TODO
	var code string
	if code, _ = cache.Get(cache.GetPhoneCodeKey(req.GetPhone())); code != req.GetAuthCode() {
		res.StatusCode = resp.CodeInvalidCaptcha
		return
	}

	// 校验密码复杂度
	if !password.CheckPassword(req.GetPassword()) {
		res.StatusCode = resp.CodeInvalidPassword
		return
	}

	// 插入数据库
	user := &model.User{
		ID:       random.GetSnowIDbyStr(),
		Phone:    req.GetPhone(),
		Avatar:   defaultAvatar,
		Password: password.Encrypt(req.GetPassword()),
		Role:     model.ConstRoleOfUser,
	}

	var e *mysql.MySQLError
	err = db.InsertUser(user)
	if errors.As(err, &e) && e.Number == 1062 {
		res.StatusCode = resp.CodeUserExist
		return
	} // 校验唯一用户
	if err != nil {
		zap.L().Error("db.InsertUser fail", zap.Error(err))
		return
	}

	res.StatusCode = resp.CodeSuccess
	return
}

// SendPhoneCode implements the AuthService interface.
func (s *AuthServiceImpl) SendPhoneCode(ctx context.Context, req *auth.SendPhoneCodeReq) (res *auth.SendPhoneCodeResp, err error) {
	res = new(auth.SendPhoneCodeResp)
	res.StatusCode = resp.CodeServerBusy

	// 校验发送间隔
	var result string
	if result, err = cache.Get(cache.GetSendCaptchaIntervalKey(req.GetPhone())); len(result) != 0 || err == nil {
		res.StatusCode = resp.CodeRateLimitExceeded
		return
	}

	// 生成code
	Captcha := random.GetRandomNum(defaultCaptchaLength)

	// 发送code
	phoneConf := config.GetConf().Phone

	if err = sms.SendCaptcha(req.GetPhone(), Captcha); err != nil {
		zap.L().Error("手机验证码发送失败" + err.Error())
		return
	}

	// 删除原来的验证码
	cache.Del(cache.GetPhoneCodeKey(req.GetPhone()))

	// 放入缓存
	cache.SetWithTime(cache.GetSendCaptchaIntervalKey(req.GetPhone()), "1", parse.Duration(phoneConf.SendInterval)) // 刷新间隔
	_, err = cache.SetWithTime(cache.GetPhoneCodeKey(req.GetPhone()), Captcha, parse.Duration(phoneConf.ExpirationTime))
	if err != nil {
		zap.L().Error("redis.Set fail", zap.Error(err))
		return
	}

	res.StatusCode = resp.CodeSuccess
	return
}

// SendEmailCode implements the AuthService interface.
const defaultCaptchaLength = 6 // TODO : 写入配置文件
func (s *AuthServiceImpl) SendEmailCode(ctx context.Context, req *auth.SendEmailCodeReq) (res *auth.SendEmailCodeResp, err error) {
	res = new(auth.SendEmailCodeResp)
	res.StatusCode = resp.CodeServerBusy

	// 校验发送间隔
	var result string
	if result, err = cache.Get(cache.GetSendCaptchaIntervalKey(req.GetEmail())); len(result) != 0 || err == nil {
		res.StatusCode = resp.CodeRateLimitExceeded
		return
	}

	// 生成code
	Captcha := random.GetRandomNum(defaultCaptchaLength)

	// 发送code
	emailConf := config.GetConf().Email

	if err = mail.SendCaptcha(req.GetEmail(), Captcha); err != nil {
		zap.L().Error("邮箱验证码发送失败" + err.Error())
		return
	}

	// 删除原来的验证码
	cache.Del(cache.GetEmailKey(req.GetEmail()))

	// 放入缓存
	cache.SetWithTime(cache.GetSendCaptchaIntervalKey(req.GetEmail()), "1", parse.Duration(config.GetConf().Email.SendInterval)) // 刷新间隔
	_, err = cache.SetWithTime(cache.GetEmailKey(req.GetEmail()), Captcha, parse.Duration(emailConf.ExpirationTime))
	if err != nil {
		zap.L().Error("redis.Set fail", zap.Error(err))
		return
	}

	res.StatusCode = resp.CodeSuccess
	return
}

// ShowPhotoCaptcha implements the AuthService interface.
func (s *AuthServiceImpl) ShowPhotoCaptcha(ctx context.Context, req *auth.ShowPhotoCaptchaReq) (res *auth.ShowPhotoCaptchaResp, err error) {
	res = new(auth.ShowPhotoCaptchaResp)
	res.StatusCode = resp.CodeServerBusy

	res.CaptchaId, res.CaptchaImg, _, err = captcha.NewCapt().GenerateCaptcha()
	if err != nil {
		zap.L().Error("生成验证码失败", zap.Error(err))
		return
	}

	res.StatusCode = resp.CodeSuccess
	return
}

// RefreshToken implements the AuthService interface.
func (s *AuthServiceImpl) RefreshToken(ctx context.Context, req *auth.RefreshTokenReq) (res *auth.RefreshTokenResp, err error) {
	res = new(auth.RefreshTokenResp)
	res.StatusCode = resp.CodeServerBusy

	// 验证并解析token
	claims, err := token.ParseToken(req.GetRefreshToken())
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			res.StatusCode = resp.CodeInvalidTokenForm
			return
		}

		if errors.Is(err, jwt.ErrTokenExpired) && claims.TokenType == 1 {
			res.StatusCode = resp.CodeInvalidTokenExpired
			return
		}

		res.StatusCode = resp.CodeInvalidToken
		return
	}

	// 读取缓存
	t, err := cache.Get(cache.GetRefreshTokenKey(claims.UserId))
	if t != req.GetRefreshToken() || err != nil {
		res.StatusCode = resp.CodeInvalidTokenExpired
		return
	}

	// 删缓存
	go cache.Del(cache.GetRefreshTokenKey(claims.ID))

	// 生成验证token
	if res.Token, err = token.GenerateAccessToken(claims.ID); err != nil || len(res.Token) == 0 {
		zap.L().Error("验证token生成失败", zap.Error(err))
		return
	}

	// 生成刷新token
	if res.RefreshToken, err = token.GenerateRefreshToken(claims.ID); err != nil || len(res.RefreshToken) == 0 {
		zap.L().Error("刷新token生成失败", zap.Error(err))
		return
	}

	// 存入缓存
	_, err = cache.SetWithTime(cache.GetRefreshTokenKey(claims.ID), res.RefreshToken, parse.Duration(config.GetConf().Jwt.RefreshExpireTime))
	if err != nil {
		zap.L().Error("redis.Set fail", zap.Error(err))
		return
	}

	// success
	res.StatusCode = resp.CodeSuccess

	return
}

// GetUserAdmin implements the AuthService interface.
const expireTime = "7d" // TODO 写入配置文件
func (s *AuthServiceImpl) GetUserAdmin(ctx context.Context, req *auth.CheckAdminReq) (res *auth.CheckAdminResp, err error) {
	res = new(auth.CheckAdminResp)
	res.StatusCode = resp.CodeServerBusy

	// 尝试命中缓存
	AdminRole, err := cache.Get(cache.GetUserRoleKey(req.GetUserId()))
	res.Role, err = strconv.ParseInt(AdminRole, 10, 64)
	if err == nil && len(AdminRole) != 0 {
		res.StatusCode = resp.CodeSuccess
		return
	}

	// 回源获取信息
	User, err := db.GetUserByID(req.GetUserId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res.StatusCode = resp.CodeRecordNotFound
			return
		}
		zap.L().Error("GetUserByID failed", zap.Error(err))
		return
	}

	// 重新存入缓存
	//expireTime := config.GetConf().RoleCacheExpireTime

	go cache.SetWithTime(cache.GetUserRoleKey(req.GetUserId()), User.Role, parse.Duration(expireTime))

	res.Role = User.Role
	res.StatusCode = resp.CodeSuccess
	return
}
