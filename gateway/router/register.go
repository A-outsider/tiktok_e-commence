package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"gomall/gateway/middleware"
)

func Register(r *server.Hertz) {
	r.Use(middleware.CORS())
}
