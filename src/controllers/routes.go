package controllers

import (
	"github.com/gofiber/fiber/v2"
	Bcontrollers "github.com/matheusgb/cyclists/src/controllers/bikeEvent"
	Ucontrollers "github.com/matheusgb/cyclists/src/controllers/user"
)

func UserRoutes(app fiber.Router, controller Ucontrollers.IUser) {
	User := app.Group("/user")
	User.Post("/", controller.CreateUser)
	User.Get("/", controller.GetAllUsers)
	User.Get("/:id", controller.GetUser)
	User.Put("/:id", controller.UpdateUser)
	User.Delete("/:id", controller.DeleteUser)
}

func BikeEventRoutes(app fiber.Router, controller Bcontrollers.IBikeEvent) {
	User := app.Group("/bike-event")
	User.Post("/", controller.CreateBikeEvent)
}
