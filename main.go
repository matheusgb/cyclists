package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheusgb/cyclists/config"
	gorm "github.com/matheusgb/cyclists/database"
	controllers "github.com/matheusgb/cyclists/src/controllers/user"
	"github.com/matheusgb/cyclists/src/models/repositories"
	"github.com/matheusgb/cyclists/src/models/services"
)

func main() {
	config := config.Init()

	database := gorm.Init()
	database.Connect(config)
	database.RunMigrations(config)

	repository := repositories.Init(database.GetClient())
	service := services.Init(repository)
	controller := controllers.Init(service)

	app := fiber.New()
	v1 := app.Group("/api/v1")
	UserRoutes(v1, controller)

	app.Listen(config.Api.Port)
}
