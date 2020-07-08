package models

import "github.com/zedisdog/armor/model"

type Permission struct {
	model.Model
	Name        string `gorm:"type:varchar(255) COMMENT '用户名';" json:"name"`
	DisplayName string `gorm:"type:varchar(255) COMMENT '显示名称';DEFAULT:''" json:"display_name"`
}

type SetPermission func(a *Permission)

func NewPermission() *Permission {
	permission := &Permission{}
	permission.SetIDWithSnowFlake()
	return permission
}
