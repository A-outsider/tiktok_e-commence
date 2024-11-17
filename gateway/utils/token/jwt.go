package token

import (
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"gomall/gateway/config"
)

// -------------------------------------jwt生成token加密------------------------------------------------

type Claims struct {
	UserId    string
	TokenType uint // 0:accessToken,1:refreshToken
	jwt.RegisteredClaims
}

func ParseToken(tokenString string) (*Claims, error) {
	// 获取jwt的荷载数据
	claims := new(Claims)
	parser := jwt.NewParser()

	_, _, err := parser.ParseUnverified(tokenString, claims) // 不验证签名获取荷载数据
	if err != nil {
		zap.L().Error("token荷载解析失败", zap.Error(err))
		return nil, err
	}

	// 判断类型 选择不同的密钥
	var secret []byte
	if claims.TokenType == 0 {
		secret = []byte(config.GetConf().Jwt.AccessSecret)
	} else if claims.TokenType == 1 {
		secret = []byte(config.GetConf().Jwt.RefreshSecret)
	}

	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		return secret, nil
	})

	return claims, err
}

func generateToken(key []byte, claims *Claims) (string, error) {
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
