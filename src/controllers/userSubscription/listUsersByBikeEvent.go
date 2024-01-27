package controllers

import (
	"github.com/gofiber/fiber/v2"
	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
)

func (UserSubscription *UserSubscription) ListUsersByBikeEvent(ctx *fiber.Ctx) error {
	bikeEventID := ctx.Params("id", "")
	if bikeEventID == "" {
		ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	domain := domains.InitID(bikeEventID)

	users, err := UserSubscription.service.ListUsersByBikeEvent(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	ctx.Status(200).JSON(fiber.Map{
		"users": users,
	})
	return nil

}
