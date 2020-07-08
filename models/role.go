package models

import (
	"github.com/zedisdog/armor/model"
	"reflect"
)

const (
	ADMIN      = "admin"
	ENTERPRISE = "enterprise"
)

type Role struct {
	model.Model
	Name        string       `gorm:"type:varchar(255) COMMENT '用户名';" json:"name"`
	DisplayName string       `gorm:"type:varchar(255) COMMENT '显示名称';DEFAULT:''" json:"display_name"`
	Permissions []Permission `gorm:"many2many:permission_role;" json:"permissions"`
}

func (r *Role) Can(name string) bool {
	for _, perm := range r.Permissions {
		if perm.Name == name {
			return true
		}
	}

	return false
}

func (r *Role) Set(field string, username interface{}) *Role {
	f := reflect.ValueOf(r).Elem().FieldByName(field)
	if f.CanSet() {
		f.Set(reflect.ValueOf(username))
	}
	return r
}

type SetRole func(r *Role)

func WithName(name string) SetRole {
	return func(r *Role) {
		r.Set("Name", name)
	}
}

func WithDisplayName(displayName string) SetRole {
	return func(r *Role) {
		r.Set("DisplayName", displayName)
	}
}

func NewRole(setFuncs ...SetRole) *Role {
	role := &Role{}
	for _, f := range setFuncs {
		f(role)
	}
	role.SetIDWithSnowFlake()
	return role
}
