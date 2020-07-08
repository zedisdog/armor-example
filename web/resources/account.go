package resources

import (
	"github.com/zedisdog/armor/helper"
	"hermes/models"
)

func Account(a models.Account) (resource map[string]interface{}) {
	resource = helper.Only(helper.Struct2Map(a), []string{
		"username",
		"avatar",
		"mobile",
		"roles",
	})
	return
}
