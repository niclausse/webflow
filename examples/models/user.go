package models

import (
	"github.com/niclausse/webflow/layer"
	"github.com/niclausse/webkit/typex"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const (
	TblUser = "tblUser"
)

type User struct {
	ID        int64          `gorm:"column:id" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	Ping      string         `gorm:"column:ping" json:"ping"`
	RealName  string         `gorm:"column:real_name" json:"realName"`
	CreatedAt typex.TimeX    `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt typex.TimeX    `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`
}

type UserDao struct {
	layer.Dao
}

func (entity *UserDao) OnCreate(param layer.IFlowParam) {
	entity.Param = param
	entity.SetTableName(TblUser)
}

func (entity *UserDao) GetByID(id int64) (user User, err error) {
	err = entity.GetDB().Where("id = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return
}
