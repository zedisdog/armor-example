package services

import (
	"errors"
	"github.com/zedisdog/armor/model"
	"hermes/models"
	"hermes/repositories"
)

func CreateEnterprise(mobile string, password string) (*models.Account, error) {
	account := models.NewAccount(
		models.WithMobile(mobile),
		models.WithPassword(password),
	)
	var role *models.Role
	if role = repositories.GetRoleByName(models.ENTERPRISE); role == nil {
		return nil, errors.New("role not found")
	}

	model.DB.Create(account).Association("roles").Append(role)

	return account, nil
}
