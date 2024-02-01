package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheusgb/cyclists/config"
	gorm "github.com/matheusgb/cyclists/database"
	routes "github.com/matheusgb/cyclists/src/controllers"
	"github.com/matheusgb/cyclists/src/controllers/validators"
)

func main() {
	// TODO: Add validations using database (unique email, etc)
	// usuario unico email, se o usuario ja esta inscrito no evento
	// se ja existe um evento com o mesmo nome no organizador
	// se o evento ja passou
	// se o evento ja esta cheio
	// talvez pra fazer essas validações de data eu tenha que iniciar o
	// postgres em um docker localmente
	// TODO: Add Logs
	// TODO: Add JWT Login
	// TODO: Add Rate Limit
	// TODO: Change password using sendgrid
	// TODO: Add Redis for cache
	// TODO: Add Kafka for webhooks (?)
	// TODO: Add Tests
	// TODO: Add Docker
	// TODO: Add CI/CD
	// TODO: Add Swagger
	config := config.Init()

	database := gorm.Init()
	database.Connect(config)
	database.RunMigrations(config)
	databaseClient := database.GetClient()

	validators.Init()

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
