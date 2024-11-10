package controller

import (
	"github.com/cloudwego/hertz/pkg/app"
	"gomall/gateway/types/resp"
	"net/http"
)

// Controller definition
type Controller[T any] struct {
	Request  *T
	Response *resp.Response
	c        *app.RequestContext
}

// NewCtrl Generic factory function for creating a controller
func NewCtrl[T any](c *app.RequestContext) *Controller[T] {
	return &Controller[T]{
		Request:  new(T),
		Response: new(resp.Response),
		c:        c,
	}
}

// NoDataJSON parse with Nodata to json and return
func (ctrl *Controller[T]) NoDataJSON(code int64) {
	ctrl.Response.SetNoData(code)
	ctrl.c.JSON(http.StatusOK, ctrl.Response)
}

// WithDataJSON parse with data to json and return
func (ctrl *Controller[T]) WithDataJSON(code int64, data interface{}) {
	ctrl.Response.SetWithData(code, data)
	ctrl.c.JSON(http.StatusOK, ctrl.Response)
}
