package responses

import (
	"hermes/models"
)

// Login login response
type Login struct {
	Token    string        `json:"token"`
	Username string        `json:"username"`
	Mobile   string        `json:"mobile"`
	Roles    []models.Role `json:"roles"`
}
