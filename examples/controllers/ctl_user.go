package controllers

import (
	"github.com/penglin1995/webflow/examples/dto"
	"github.com/penglin1995/webflow/examples/logic"
	"github.com/penglin1995/webflow/layer"
)

type AddUserCTL struct {
	layer.Controller
}

func (c *AddUserCTL) SetBindingObject() {
	c.DTO = new(dto.AddUserReq)
}

func (c *AddUserCTL) Action() (interface{}, error) {
	return nil, logic.NewUserLogic().Add(c.GetContext(), c.GetBindingObject().(*dto.AddUserReq))
}
