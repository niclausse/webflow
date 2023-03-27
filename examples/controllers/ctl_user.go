package controllers

import (
	"github.com/niclausse/webflow/examples/service"
	"github.com/niclausse/webflow/layer"
)

type UserInfoReq struct {
	Uid int64 `form:"uid" binding:"required"`
}

type UserInfoCTL struct {
	layer.Controller
	dto *UserInfoReq
}

func (entity *UserInfoCTL) GetBindingObject() interface{} {
	return &entity.dto
}

func (entity *UserInfoCTL) Action() (interface{}, error) {
	return entity.Create(new(service.UserService)).(*service.UserService).UserInfo(entity.dto.Uid)
}
