package service

import (
	"github.com/niclausse/webflow/examples/models"
	"github.com/niclausse/webflow/layer"
)

type UserService struct {
	layer.Service
}

func (entity *UserService) UserInfo(uid int64) (models.User, error) {
	return entity.Create(new(models.UserDao)).(*models.UserDao).GetByID(uid)
}
