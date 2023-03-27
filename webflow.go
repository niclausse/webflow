package webflow

import (
	"github.com/gin-gonic/gin"
	"github.com/niclausse/webflow/layer"
	"github.com/niclausse/webkit/errorx"
	"github.com/niclausse/webkit/response"
	"gorm.io/gorm"
)

func UseController(ctl layer.IController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctl.SetGinContext(ctx)

		if req := ctl.GetBindingObject(); req != nil {
			if bindingType := ctl.GetBindingType(); bindingType != nil {
				if err := ctx.ShouldBindWith(req, bindingType); err != nil {
					response.Fail(ctx, errorx.ParamInvalid.WithDetails(err.Error()).WithStack())
					return
				}
			}

			if err := ctx.ShouldBind(req); err != nil {
				response.Fail(ctx, errorx.ParamInvalid.WithDetails(err.Error()).WithStack())
			}
		}

		resp, err := ctl.Action()
		if err != nil {
			response.Fail(ctx, err)
			return
		}

		response.Succeed(ctx, resp)
	}
}

func SetDefaultDB(db *gorm.DB) {
	layer.SetDefaultDB(db)
}
