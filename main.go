package main

import (
	"github.com/gofiber/fiber/v2"
	routes "github.com/matheusgb/cyclists/api"
)

func main() {
	app := fiber.New()
	v1 := app.Group("/api/v1")
	routes.BikeEvents(v1)

	app.Listen(":3000")
}
