package main

import (
	"github.com/zedisdog/armor"
	"github.com/zedisdog/armor/log"
	"hermes/models"
	"hermes/web"
)

//go:generate cp .env.example .env
func main() {
	server := armor.NewArmor(models.AutoMigrate)
	models.InitSeeder()
	models.TestSeeder()
	err := server.Start(web.MakeRoutes)
	if err != nil {
		log.Log.WithError(err).Info("server start failed")
	}
}
