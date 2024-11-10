package model

import (
	"github.com/cloudwego/hertz/pkg/route"
)

func RegisterUser(r *route.RouterGroup) {
	u := r.Group("/user")
	u.GET("ping")
}
