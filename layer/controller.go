package layer

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type IController interface {
	IFlow
	SetGinContext(ctx *gin.Context)
	GetGinContext() *gin.Context
	GetBindingType() binding.Binding
	GetBindingObject() interface{}
	Action() (interface{}, error)
}

type Controller struct {
	Flow
	gx *gin.Context
}

func (d *Controller) SetGinContext(ctx *gin.Context) {
	d.gx = ctx
	d.SetContext(ctx.Request.Context())
}

func (d *Controller) GetGinContext() *gin.Context {
	return d.gx
}

func (d *Controller) GetBindingType() binding.Binding {
	return nil
}

func (d *Controller) GetBindingObject() interface{} {
	return nil
}

func (d *Controller) Action() (interface{}, error) {
	return "implement me", nil
}
