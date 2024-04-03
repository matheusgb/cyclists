package controllers

import (
	"github.com/gofiber/fiber/v2"
	Bcontrollers "github.com/matheusgb/cyclists/src/controllers/bikeEvent"
	Ucontrollers "github.com/matheusgb/cyclists/src/controllers/user"
	UScontrollers "github.com/matheusgb/cyclists/src/controllers/userSubscription"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
)

func UserRoutes(app fiber.Router, controller Ucontrollers.IUser) {
	User := app.Group("/user")
	User.Use(domains.VerifyJWTTokenMiddleware)
	User.Get("/", controller.GetAllUsers)
	User.Get("/:id", controller.GetUser)
	User.Put("/:id", controller.UpdateUser)
	User.Delete("/:id", controller.DeleteUser)

	Login := app.Group("/login")
	Login.Post("/", controller.LoginUser)

	Register := app.Group("/register")
	Register.Post("/", controller.CreateUser)
}

func BikeEventRoutes(app fiber.Router, controller Bcontrollers.IBikeEvent) {
	bikeEvent := app.Group("/bike-event")
	bikeEvent.Use(domains.VerifyJWTTokenMiddleware)
	bikeEvent.Post("/", controller.CreateBikeEvent)
	bikeEvent.Get("/", controller.GetAllBikeEvents)
	bikeEvent.Get("/:id", controller.GetBikeEvent)
	bikeEvent.Patch("/:id", controller.UpdateBikeEvent)
	bikeEvent.Delete("/:id", controller.DeleteBikeEvent)
}

func UserSubscriptionRoutes(app fiber.Router, controller UScontrollers.IUserSubscription) {
	userSubscription := app.Group("/user-subscription")
	userSubscription.Use(domains.VerifyJWTTokenMiddleware)
	userSubscription.Post("/", controller.CreateUserSubscription)
	userSubscription.Delete("/", controller.DeleteUserSubscription)
}
