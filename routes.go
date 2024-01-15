package main

import "github.com/gofiber/fiber/v2"

func UserRoutes(app fiber.Router) {
	User := app.Group("/user")
	User.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
