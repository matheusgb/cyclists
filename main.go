package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheusgb/cyclists/src/config"
	routes "github.com/matheusgb/cyclists/src/controllers"
	"github.com/matheusgb/cyclists/src/controllers/validators"
	gorm "github.com/matheusgb/cyclists/src/database"
	"github.com/matheusgb/cyclists/src/layers"
)

func main() {
	// TODO: Add JWT Login and create context for user
	// TODO: Add Logs
	// TODO: Add Redis for cache
	// TODO: Add Tests
	// TODO: Add Swagger
	// TODO: Add Docker
	// TODO: Add CI/CD

	//! v2
	// TODO: Change password using sendgrid
	// TODO: Add date validation for bike event (if a event is in the past, it should not be possible to subscribe)
	// TODO: Add Rate Limit
	// TODO: Add Kafka for webhooks

	config := config.Init()

	database := gorm.Init()
	database.Connect(config)
	database.RunMigrations(config)
	databaseClient := database.GetClient()

	validators.Init()

	userControllers := layers.InitUser(databaseClient)
	bikeEventControllers := layers.InitBikeEvent(databaseClient)
	userSubscriptionControllers := layers.InitUserSubscription(databaseClient)

	app := fiber.New()
	v1 := app.Group("/api/v1")
	routes.UserRoutes(v1, userControllers)
	routes.BikeEventRoutes(v1, bikeEventControllers)
	routes.UserSubscriptionRoutes(v1, userSubscriptionControllers)

	app.Listen(config.Api.Port)
}
