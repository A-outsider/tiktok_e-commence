package token

import (
	"dream_program/util/parse"
	"time"

	"dream_program/config"

	"github.com/golang-jwt/jwt/v5"
)

/**
 * 生成验证用的token
 * param: id 用户id
 * return: token字符串、错误信息
 */
func GenerateAccessToken(id int64) (string, error) {
	accessSecret := []byte(config.Get().Auth.AccessJwtSecret)
	Issuer := config.Get().Auth.Issuer
	expireTime := config.Get().Auth.AccessJwtExpireTime

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
func GenerateRefreshToken(id int64) (string, error) {
	refreshSecret := []byte(config.Get().Auth.RefreshJwtSecret)
	Issuer := config.Get().Auth.Issuer
	expireTime := config.Get().Auth.RefreshJwtExpireTime

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
