package controllers

import (
	"github.com/gofiber/fiber/v2"
	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
)

func (UserSubscription *UserSubscription) DeleteUserSubscription(ctx *fiber.Ctx) error {
	userSubscriptionID := ctx.Params("id", "")
	if userSubscriptionID == "" {
		ctx.Status(400).JSON(fiber.Map{
			"message": "userSubscription ID is required",
		})
		return nil
	}

	domain := domains.InitID(userSubscriptionID)

	_, err := UserSubscription.service.DeleteUserSubscription(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	ctx.Status(204)
	return nil
}
