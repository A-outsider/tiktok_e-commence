package auth

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type Api struct{}

func NewApi() *Api {
	return &Api{}
}

func (api *Api) Login(ctx context.Context, c *app.RequestContext) {

}
