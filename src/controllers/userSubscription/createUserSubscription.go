package controllers

import (
	"github.com/gofiber/fiber/v2"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/userSubscription"
	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	views "github.com/matheusgb/cyclists/src/views/userSubscription"
)

func (UserSubscription *UserSubscription) CreateUserSubscription(ctx *fiber.Ctx) error {
	var request requests.CreateUserSubscription

	if err := ctx.BodyParser(&request); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
		return err
	}

	domain := domains.InitCreate(request.BikeEventID, request.UserID)

	entity, err := UserSubscription.service.CreateUserSubscription(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	view := views.ConvertUserSubscriptionEntityToResponse(&entity)
	ctx.Status(201).JSON(view)
	return nil
}
