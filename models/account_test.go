package models

import (
	"fmt"
	"testing"
)

func TestAccount(t *testing.T) {
	a := &Account{
		Roles: []Role{
			{
				Name: "test",
				Permissions: []Permission{
					{
						Name: "create-post",
					},
				},
			},
		},
	}

	if a.Can("create-post") != true {
		t.Error(fmt.Sprintf("except true, got %t", a.Can("create-post")))
	}
}

func TestAccount_Set(t *testing.T) {
	account := Account{}
	account.Username = "admin"
	if account.Username != "admin" {
		t.Error(fmt.Sprintf("except admin, got <%s>", account.Username))
	}

	id := uint64(32132132131)
	account.ID = id
	if account.ID != id {
		t.Error(fmt.Sprintf("except ID is 32132132131 with uint64, got <%d>", account.ID))
	}
}
