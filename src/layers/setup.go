package layers

import (
	"github.com/gofiber/fiber/v2"
	swagger "github.com/swaggo/fiber-swagger"

	_ "github.com/matheusgb/cyclists/docs"
	"github.com/matheusgb/cyclists/src/config"
	routes "github.com/matheusgb/cyclists/src/controllers"
	"github.com/matheusgb/cyclists/src/controllers/validators"
	gorm "github.com/matheusgb/cyclists/src/database"
)

func Setup(config *config.Config) *fiber.App {
	database := gorm.Init()
	database.Connect(*config)
	database.RunMigrations()
	databaseClient := database.GetClient()

	validators.Init()

	userControllers := InitUser(databaseClient)
	bikeEventControllers := InitBikeEvent(databaseClient)
	userSubscriptionControllers := InitUserSubscription(databaseClient)

	app := fiber.New()

	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"status": "ok"})
	})

	app.Get("/docs/*", swagger.WrapHandler)

	v1 := app.Group("/api/v1")
	routes.UserRoutes(v1, userControllers)
	routes.BikeEventRoutes(v1, bikeEventControllers)
	routes.UserSubscriptionRoutes(v1, userSubscriptionControllers)

	return app
}
