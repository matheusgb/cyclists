package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheusgb/cyclists/config"
	"github.com/matheusgb/cyclists/gorm"
)

func main() {
	config := config.Init()
	client := gorm.Database().VerifyConnection(config)
	gorm.RunMigrations(client, config)

	app := fiber.New()
	v1 := app.Group("/api/v1")
	BikeEvents(v1)

	app.Listen(config.Api.Port)
}
