package controllers

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/matheusgb/cyclists/src/controllers/user"
)

func UserRoutes(app fiber.Router, controller controllers.IUser) {
	User := app.Group("/user")
	User.Post("/", controller.CreateUser)
	User.Get("/", controller.GetAllUsers)
	User.Get("/:id", controller.GetUser)
	User.Put("/:id", controller.UpdateUser)
	User.Delete("/:id", controller.DeleteUser)
}
