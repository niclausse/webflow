package webflow

import (
	"github.com/gin-gonic/gin"
	"github.com/penglin1995/webflow/errorx"
	"github.com/penglin1995/webflow/layer"
	"github.com/penglin1995/webflow/logx"
	"github.com/penglin1995/webflow/response"
	"github.com/pkg/errors"
)

var logger logx.Logger

func Init(log logx.Logger) {
	logger = log
}

func UseController(ctl layer.IController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctl.SetContext(ctx)
		ctl.SetBindingObject()

		r := response.NewResponder(response.ModeDev, logger)

		bindingType := ctl.GetBindingType()
		req := ctl.GetBindingObject()
		if req == nil {
			r.Fail(ctx, errors.WithStack(errorx.ParamInvalid.WithDetails("controller dto missed")))
			return
		}

		if err := ctx.ShouldBindWith(req, bindingType); err != nil {
			r.Fail(ctx, errors.WithStack(errorx.ParamInvalid.WithDetails(err.Error())))
			return
		}

		resp, err := ctl.Action()
		if err != nil {
			r.Fail(ctx, err)
			return
		}

		r.Succeed(ctx, resp)
	}
}
