package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func CORS() app.HandlerFunc {

	return func(c context.Context, ctx *app.RequestContext) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Max-Age", "86400")
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,DELETE")
		ctx.Header("Access-Control-Allow-Headers", "authorization,Authorization,DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		if string(ctx.Method()) == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
		} else {
			ctx.Next(c)
		}
	}
}
