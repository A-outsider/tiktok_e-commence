package middleware

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
	"gomall/gateway/types/resp/common"
	"gomall/gateway/utils/role"
	"gomall/gateway/utils/token"
	"net/http"
	"strings"
)

// Parse 宽松认证
func Parse() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		tokenString, _ := getToken(ctx)

		// 解析并校验Token
		claims, _ := token.ParseToken(tokenString)
		if len(claims.UserId) != 0 {
			ctx.Set("userId", claims.UserId)
		}
		ctx.Next(c)
	}
}

// 验证用户是否登录的中间件 -- 双token
func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		res := new(common.Response)
		res.SetNoData(common.CodeSuccess)

		// 读取验证token
		tokenString, ok := getToken(c)
		if !ok {
			res.SetNoData(common.CodeNotLogin)
			c.JSON(http.StatusOK, res)
			c.Abort()
			return
		}

		// 校验token信息
		claims, err := token.ParseToken(tokenString)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenMalformed) {
				res.SetNoData(common.CodeInvalidTokenForm)
				c.JSON(http.StatusOK, res)
				c.Abort()
				return
			}

			// 提示需要刷新token
			if errors.Is(err, jwt.ErrTokenExpired) && claims.TokenType == 0 {
				res.SetNoData(common.CodeInvalidTokenExpired)
				c.JSON(http.StatusOK, res)
				c.Abort()
				return
			}

			res.SetNoData(common.CodeInvalidToken)
			c.JSON(http.StatusOK, res)
			c.Abort()
			return
		}

		// 认证用户角色权限
		StatusCode := role.CheckAdmin(ctx, c, claims.UserId)
		if StatusCode != common.CodeSuccess {
			res.SetNoData(StatusCode)
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		// 存储用户信息
		c.Set("userId", claims.UserId)
		c.Next(ctx)
	}
}

func getToken(c *app.RequestContext) (string, bool) {
	tokenString := c.GetHeader("Authorization")

	if !strings.HasPrefix(string(tokenString), "Bearer ") {
		return "", false
	}

	tokenString = []byte(strings.TrimPrefix(string(tokenString), "Bearer "))
	return string(tokenString), true
}
