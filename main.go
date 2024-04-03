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
	// TODO: Change user password using sendgrid
	// TODO: Add Tests
	// TODO: Add Swagger

	//! v2
	// TODO: Add Docker
	// TODO: Add CI/CD
	// TODO: Add Redis for cache
	// TODO: Add Rate Limit
	// TODO: Add Kafka for webhooks
	// TODO: Add Logs and save in a file

	//? v3
	// TODO: Add Prometheus + Grafana for monitoring

	config := config.Init()
	config.MountConfigs()

	database := gorm.Init()
	database.Connect(*config)
	database.RunMigrations(*config)
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
