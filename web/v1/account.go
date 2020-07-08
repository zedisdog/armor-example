package v1

import (
	"hermes/models"
	"hermes/web/requests"
	"hermes/web/responses"

	"github.com/gin-gonic/gin"
	"github.com/zedisdog/armor/auth"
	"github.com/zedisdog/armor/model"
	"github.com/zedisdog/armor/tools"
	"github.com/zedisdog/armor/web"
)

// Login 登录api
func Login(c *gin.Context) {
	var req requests.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		_ = c.AbortWithError(400, err)
		return
	}

	var account models.Account
	not := model.DB.Where("username=?", req.Username).First(&account).RecordNotFound()

	if not {
		_ = c.AbortWithError(404, web.NotFound)
		return
	}

	result, _ := tools.Hash(req.Password)

	if result != account.Password {
		_ = c.AbortWithError(401, web.UsernameOrPasswordError)
		return
	}

	var token string
	var err error
	if token, err = auth.GenerateToken(&account); err != nil {
		_ = c.AbortWithError(500, web.ServerError)
		return
	}

	response := responses.Login{
		Token:    token,
		Username: account.Username,
		Mobile:   account.Mobile,
		Roles:    account.Roles,
	}

	c.JSON(200, response)
}

// Account 获取当前登录用户信息api
func Account(c *gin.Context) {
	var account models.Account
	if err := model.DB.Where("id=?", c.Keys["id"].(uint64)).First(&account).RecordNotFound(); err {
		_ = c.AbortWithError(404, web.NotFound)
	}
	c.JSON(200, account)
}
