package response

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/penglin1995/webflow/errorx"
	"github.com/penglin1995/webflow/logx"
)

type Render struct {
	ErrNo  int         `json:"err_no"`
	ErrMSG string      `json:"err_msg"`
	Data   interface{} `json:"data"`
}

type Responder interface {
	Fail(ctx *gin.Context, err error)
	Succeed(ctx *gin.Context, data interface{})
}

func NewResponder(mode string, logger logx.Logger) Responder {
	return &responder{runMode: mode, logger: logger}
}

type responder struct {
	runMode string
	logger  logx.Logger
}

func (r *responder) Fail(ctx *gin.Context, err error) {
	no, msg, details := errorx.Biz(err)

	resp := gin.H{
		"err_no":  no,
		"err_msg": msg,
	}

	stack := strings.Split(fmt.Sprintf("%+v", err), "\n")

	if r.runMode == "dev" {
		resp["details"] = details
		resp["stack"] = stack
	}

	r.logger.Errorf("%+v", err)

	ctx.JSON(http.StatusOK, resp)
}

func (r *responder) Succeed(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"err_no":  0,
		"err_msg": "",
		"data":    data,
	})
}
