package web

import (
	"github.com/gin-gonic/gin"
	"github.com/zedisdog/armor/auth"
	"github.com/zedisdog/armor/web"
	v1 "hermes/web/v1"
)

func MakeRoutes(r *gin.Engine) {
	r.Use(web.Cros)
	api := r.Group("/api")
	{
		api.POST("/register", v1.Register)
		api.POST("/login", v1.Login)
		api.Use(auth.Auth)
		api.GET("/account", v1.Account)
	}
	r.GET("test", v1.Test)
}
