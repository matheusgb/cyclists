package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/userSubscription"
	"github.com/matheusgb/cyclists/src/controllers/validators"
	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
)

func (UserSubscription *UserSubscription) CreateUserSubscription(ctx *fiber.Ctx) error {
	var request requests.UserSubscription

	if err := ctx.BodyParser(&request); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	errorsValidation := validators.UserSubscription(request)
	if errorsValidation != nil {
		ctx.Status(400).JSON(fiber.Map{
			"errors": errorsValidation,
		})
		return nil
	}

	userRole := ctx.Locals("user_role").(string)
	uintUserId, err := strconv.ParseUint(ctx.Locals("user_id").(string), 10, 64)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"message": "internal error",
		})
	}
	userId := uint(uintUserId)

	if request.UserID != userId && userRole != "admin" {
		ctx.Status(403).JSON(fiber.Map{
			"message": "you don't have permission to create this user subscription on behalf of another user",
		})
		return nil
	}

	domain := domains.Init(request.BikeEventID, request.UserID)

	_, err = UserSubscription.service.CreateUserSubscription(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	ctx.Status(201)
	return nil
}
