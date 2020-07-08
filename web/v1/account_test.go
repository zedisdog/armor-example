package v1

import (
	"encoding/json"
	"hermes/models"
	"hermes/web/responses"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/zedisdog/armor/model"
	"github.com/zedisdog/armor/tools"
	"github.com/zedisdog/armor/tools/test"
)

func TestAccount(t *testing.T) {
	convey.FocusConvey("test account controller", t, func() {
		db := model.DB
		model.DB = model.DB.Begin()
		models.AutoMigrate(db)

		convey.Convey("test login", func() {
			password, _ := tools.Hash("secret")
			account := models.Account{
				Username: "admin",
				Password: password,
			}
			model.DB.Create(&account)

			res, _ := test.Post(Login, test.Data{
				"username": "admin",
				"password": "secret",
			})

			convey.So(res.Code, convey.ShouldEqual, 200)
			var login responses.Login
			err := json.Unmarshal(res.Body.Bytes(), &login)
			convey.So(err, convey.ShouldBeNil)
			convey.So(login.Mobile, convey.ShouldBeEmpty)
			convey.So(login.Username, convey.ShouldEqual, "admin")
			convey.So(login.Token, convey.ShouldNotBeEmpty)
		})

		convey.Convey("test account", func() {
			password, _ := tools.Hash("secret")
			account := models.Account{
				Username: "admin",
				Password: password,
			}
			model.DB.Create(&account)

			test.Act(&account)
			res, _ := test.Get(Account)

			convey.So(res.Code, convey.ShouldEqual, 200)
			convey.So(len(res.Body.Bytes()), convey.ShouldNotEqual, 0)
		})

		convey.Reset(func() {
			model.DB.Rollback()
			model.DB = db
		})
	})
}
