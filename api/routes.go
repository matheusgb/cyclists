package routes

import "github.com/gofiber/fiber/v2"

func BikeEvents(app fiber.Router) {
	bikeEvents := app.Group("/bike-events")
	bikeEvents.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
