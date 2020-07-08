package repositories

import (
	"github.com/zedisdog/armor/model"
	"hermes/models"
)

func GetRoleByName(name string) *models.Role {
	role := &models.Role{}
	if model.DB.Where("name=?", name).First(role).RecordNotFound() {
		return nil
	} else {
		return role
	}
}
