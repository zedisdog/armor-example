package models

import (
	"errors"
	"github.com/zedisdog/armor/model"
	"strconv"
)

// InitSeeder 初始化填充
func InitSeeder() {
	roles := []map[string]string{
		{
			"name":         ADMIN,
			"display_name": "管理员",
		},
		{
			"name":         ENTERPRISE,
			"display_name": "商家",
		},
	}

	for _, r := range roles {
		if model.DB.Where("name=?", r["name"]).First(&Role{}).RecordNotFound() {
			seedRole(r["name"], r["display_name"])
		}
	}

	if model.DB.Where("username=?", "admin").First(&Account{}).RecordNotFound() {
		seedAdmin("admin", "secret")
	}
}

// TestSeeder 测试填充
func TestSeeder() {
	var count int
	model.DB.Model(&Account{}).Count(&count)
	if count <= 1 {
		for i := 0; i < 31; i++ {
			code := "admin" + strconv.Itoa(i)
			seedAdmin(code, "secret")
		}
	}
}

func seedAdmin(username string, password string) {
	account := NewAccount(
		WithUsername(username),
		WithPassword(password),
	)
	var role Role
	if model.DB.Where("name=?", ADMIN).First(&role).RecordNotFound() {
		panic(errors.New("no admin role found"))
	}
	model.DB.Create(account).Association("roles").Append(role)
}

func seedEnterprise(username string, mobile string, password string) {
	account := NewAccount(
		WithUsername(username),
		WithMobile(mobile),
		WithPassword(password),
	)
	var role Role
	if model.DB.Where("name=?", ENTERPRISE).First(&role).RecordNotFound() {
		panic(errors.New("no admin role found"))
	}
	model.DB.Create(account).Association("roles").Append(role)
}

func seedRole(name string, displayName string) {
	role := NewRole(
		WithName(name),
		WithDisplayName(displayName),
	)
	model.DB.Create(role)
}
