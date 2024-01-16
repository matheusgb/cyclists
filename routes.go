package main

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/matheusgb/cyclists/src/controllers/user"
)

func UserRoutes(app fiber.Router, controller controllers.IUser) {
	User := app.Group("/user")
	User.Post("/", controller.CreateUser)
	User.Put("/:id", controller.UpdateUser)
}
