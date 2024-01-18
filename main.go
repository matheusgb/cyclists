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
	// TODO: Add remaining CRUD (BikeEvents and eventUserSubscription)
	// TODO: Improve error handling
	// TODO: Add JWT
	// TODO: Add Logger
	// TODO: Change password using sendgrid
	// TODO: Add Redis
	// TODO: Add Tests
	// TODO: Add Docker
	// TODO: Add CI/CD
	// TODO: Add Swagger
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
