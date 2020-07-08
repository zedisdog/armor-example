package models

import (
	"github.com/zedisdog/armor/model"
	"github.com/zedisdog/armor/tools"
)

type Account struct {
	model.Model
	Username   string `gorm:"type:varchar(255) COMMENT '用户名';default:''" json:"username"`
	Password   string `gorm:"type:varchar(255) COMMENT '密码';default:''" json:"-"`
	Mobile     string `gorm:"type:varchar(255) COMMENT '手机号';default:''" json:"mobile"`
	MiniOpenid string `gorm:"type:varchar(255) COMMENT '小程序openid';default:''" json:"mini_openid"`
	Avatar     string `gorm:"type:varchar(255) COMMENT '头像链接';default:''" json:"avatar"`
	SessionKey string `gorm:"type:varchar(255) COMMENT '小程序session_key'" json:"-"`
	Roles      []Role `gorm:"many2many:role_user;" json:"roles"`
}

//func (a *Account) Set(field string, value interface{}) *Account {
//	f := reflect.ValueOf(a).Elem().FieldByName(field)
//	if f.CanSet() {
//		f.Set(reflect.ValueOf(value))
//	}
//	return a
//}

type SetAccount func(a *Account)

func WithUsername(username string) SetAccount {
	return func(a *Account) {
		a.Username = username
	}
}

func WithMobile(mobile string) SetAccount {
	return func(a *Account) {
		a.Mobile = mobile
	}
}

func WithPassword(password string) SetAccount {
	hash, err := tools.Hash(password)
	if err != nil {
		panic(err)
	}
	return func(a *Account) {
		a.Password = hash
	}
}

func NewAccount(withFunc ...SetAccount) *Account {
	account := &Account{}
	for _, f := range withFunc {
		f(account)
	}
	account.SetIDWithSnowFlake()
	return account
}

func (a *Account) HasRole(name string) bool {
	for _, role := range a.Roles {
		if role.Name == name {
			return true
		}
	}

	return false
}

func (a *Account) Can(name string) bool {
	for _, role := range a.Roles {
		if role.Can(name) {
			return true
		}
	}

	return false
}
