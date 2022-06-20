package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/penglin1995/webflow"
	"github.com/penglin1995/webflow/errorx"
	"github.com/penglin1995/webflow/examples/dto"
	"github.com/penglin1995/webflow/layer"
	"github.com/penglin1995/webflow/logx"
	"github.com/pkg/errors"
)

func main() {
	webflow.Init(logx.New(logx.LevelError, "./"))

	e := gin.New()

	e.POST("/ok", webflow.UseController(new(OkCTL)))
	e.GET("/fail", webflow.UseController(new(FailCTL)))

	e.Run(":8090")
}

type OkReq struct {
}

type OkCTL struct {
	req *OkReq
	layer.IController
}

func (ctl *OkCTL) GetBindingType() binding.Binding {
	return binding.JSON
}

func (ctl *OkCTL) GetRequest() interface{} {
	return &OkCTL{}
}

func (ctl *OkCTL) Action() (interface{}, error) {
	return gin.H{"ping": "pong"}, nil
}

type FailCTL struct {
	//layer.IController
}

func (f *FailCTL) GetBindingType() binding.Binding {
	return binding.Query
}

func (f *FailCTL) GetRequest() interface{} {
	return dto.FailRequest{}
}

func (f *FailCTL) Action() (interface{}, error) {
	return nil, errors.WithStack(errorx.SystemError)
}
