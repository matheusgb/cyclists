package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheusgb/cyclists/config"
	gorm "github.com/matheusgb/cyclists/database"
	routes "github.com/matheusgb/cyclists/src/controllers"
)

func main() {
	// TODO: Add remaining CRUD (eventUserSubscription)
	// TODO: Improve error handling
	// TODO: Add validations (email, password, etc)
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
	databaseClient := database.GetClient()

	userControllers := InitUserLayers(databaseClient)
	bikeEventControllers := InitBikeEventLayers(databaseClient)
	userSubscriptionControllers := InitUserSubscriptionLayers(databaseClient)

	app := fiber.New()
	v1 := app.Group("/api/v1")
	routes.UserRoutes(v1, userControllers)
	routes.BikeEventRoutes(v1, bikeEventControllers)
	routes.UserSubscriptionRoutes(v1, userSubscriptionControllers)

	app.Listen(config.Api.Port)
}
