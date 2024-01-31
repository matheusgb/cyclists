package controllers

import (
	"github.com/gofiber/fiber/v2"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/userSubscription"
	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	"github.com/matheusgb/cyclists/validator"
)

func (UserSubscription *UserSubscription) CreateUserSubscription(ctx *fiber.Ctx) error {
	var request requests.CreateUserSubscription

	if err := ctx.BodyParser(&request); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	errorsValidation := validator.UserSubscription(request)
	if errorsValidation != nil {
		ctx.Status(400).JSON(fiber.Map{
			"errors": errorsValidation,
		})
		return nil
	}

	domain := domains.InitCreate(request.BikeEventID, request.UserID)

	_, err := UserSubscription.service.CreateUserSubscription(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	ctx.Status(201)
	return nil
}
