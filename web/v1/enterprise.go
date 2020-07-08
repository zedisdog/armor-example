package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/zedisdog/armor/web"
	"hermes/services"
	"hermes/web/requests"
	"hermes/web/resources"
)

// Register 商家入驻
func Register(c *gin.Context) {
	var request requests.Register
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(422, web.ParseValidateErrors(err.(validator.ValidationErrors)))
		return
	}
	if account, err := services.CreateEnterprise(request.Mobile, request.Password); err != nil {
		c.JSON(404, gin.H{
			"message": "role not found",
		})
	} else {
		c.JSON(201, resources.Account(*account))
	}
}
