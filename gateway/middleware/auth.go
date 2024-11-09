package middleware

import (
	"errors"
	"gomall/gateway/types/resp"
	"gomall/services/auth/utils/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Parse 宽松认证
func Parse() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, _ := getToken(c)

		// 解析并校验Token
		claims, _ := token.ParseToken(tokenString)
		if claims.UserId != 0 {
			c.Set("userId", claims.UserId)
		}
		c.Next()
	}
}

// 验证用户是否登录的中间件 -- 双token
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := new(resp.Response)
		res.SetNoData(resp.CodeSuccess)

		// 读取验证token
		tokenString, ok := getToken(c)
		if !ok {
			res.SetNoData(resp.CodeNotLogin)
			c.JSON(http.StatusOK, res)
			c.Abort()
			return
		}

		// 校验token信息
		claims, err := token.ParseToken(tokenString)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenMalformed) {
				res.SetNoData(resp.CodeInvalidTokenForm)
				c.JSON(http.StatusOK, res)
				c.Abort()
				return
			}

			// 提示需要刷新token
			if errors.Is(err, jwt.ErrTokenExpired) && claims.TokenType == 0 {
				res.SetNoData(resp.CodeInvalidTokenExpired)
				c.JSON(http.StatusOK, res)
				c.Abort()
				return
			}

			res.SetNoData(resp.CodeInvalidToken)
			c.JSON(http.StatusOK, res)
			c.Abort()
			return
		}
		// 存储用户信息
		c.Set("userId", claims.UserId)
		c.Next()
	}
}

func getToken(c *gin.Context) (string, bool) {
	tokenString := c.GetHeader("Authorization")

	if !strings.HasPrefix(tokenString, "Bearer ") {
		return "", false
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	return tokenString, true
}
