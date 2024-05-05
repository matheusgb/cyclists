package controllers

import (
	"github.com/gofiber/fiber/v2"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/user"
	"github.com/matheusgb/cyclists/src/controllers/validators"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
)

func (user *User) ResetPassword(ctx *fiber.Ctx) error {
	var request requests.ResetPassword

	if err := ctx.BodyParser(&request); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	errValidator := validators.User(request)
	if errValidator != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": errValidator,
		})
		return nil
	}

	contextEmail := ctx.Locals("user_email").(string)
	if contextEmail != request.Email {
		ctx.Status(401).JSON(fiber.Map{
			"message": "invalid token",
		})
		return nil
	}

	contextUserID := ctx.Locals("user_id").(string)
	domain := domains.InitResetPassword(contextUserID, request.Password)

	err := user.service.ResetPassword(*domain)
	if err != nil {
		ctx.Status(401).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	ctx.Status(200).JSON(fiber.Map{
		"message": "password changed",
	})
	return nil
}
