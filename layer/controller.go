package layer

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type IController interface {
	SetContext(ctx *gin.Context)
	GetContext() *gin.Context
	GetBindingType() binding.Binding
	GetBindingObject() interface{}
	SetBindingObject()
	Action() (interface{}, error)
}

type Controller struct {
	gx  *gin.Context
	DTO interface{}
}

type FlowParam struct {
}

func (d *Controller) SetContext(ctx *gin.Context) {
	d.gx = ctx
}

func (d *Controller) GetContext() *gin.Context {
	return d.gx
}

func (d *Controller) GetBindingType() binding.Binding {
	return binding.JSON
}

func (d *Controller) SetBindingObject() {
	d.DTO = &FlowParam{}
}

func (d *Controller) GetBindingObject() interface{} {
	return d.DTO
}

func (d *Controller) Action() (interface{}, error) {
	return "implement me", nil
}
