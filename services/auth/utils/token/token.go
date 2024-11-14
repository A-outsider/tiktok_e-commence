package token

import (
	"gomall/common/utils/parse"
	"gomall/gateway/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/**
* 生成验证用的token
* param: id 用户id
* return: token字符串、错误信息
 */
func GenerateAccessToken(id string) (string, error) {
	accessSecret := []byte(config.GetConf().Jwt.AccessSecret)
	Issuer := config.GetConf().Jwt.Issuer
	expireTime := config.GetConf().Jwt.AccessExpireTime

	// token过期时间
	expirationTime := time.Now().Add(parse.Duration(expireTime))
	accessClaims := &Claims{
		UserId:    id,
		TokenType: 0,
		RegisteredClaims: jwt.RegisteredClaims{
			// 发放时间等
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    Issuer,
		},
	}

	return generateToken(accessSecret, accessClaims)
}

/**
* 生成刷新用的token
* param: id 用户id
* return: token字符串、错误信息
 */
func GenerateRefreshToken(id string) (string, error) {
	refreshSecret := []byte(config.GetConf().Jwt.AccessSecret)
	Issuer := config.GetConf().Jwt.Issuer
	expireTime := config.GetConf().Jwt.RefreshExpireTime

	// token过期时间
	expirationTime := time.Now().Add(parse.Duration(expireTime)) // 7天有效
	refreshClaims := &Claims{
		UserId:    id,
		TokenType: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			// 发放时间等
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    Issuer,
		},
	}

	return generateToken(refreshSecret, refreshClaims)
}
