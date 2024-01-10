package main

import (
	"github.com/gofiber/fiber/v2"
	routes "github.com/matheusgb/cyclists/api"
	"github.com/matheusgb/cyclists/config"
	"github.com/matheusgb/cyclists/database"
)

func main() {
	config := config.Init()
	database.CreateGormDatabase().VerifyConnection(config)

	app := fiber.New()
	v1 := app.Group("/api/v1")
	routes.BikeEvents(v1)

	app.Listen(config.Api.Port)
}
