package controllers

import (
	"github.com/gofiber/fiber/v2"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/user"
	"github.com/matheusgb/cyclists/src/controllers/validators"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
)

func (user *User) SendPasswordResetEmail(ctx *fiber.Ctx) error {
	var request requests.SendPasswordResetEmail

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

	domain := domains.InitSendPasswordResetEmail(request.Email)

	err := user.service.SendPasswordResetEmail(*domain)
	if err != nil {
		ctx.Status(401).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}
	ctx.Status(200).JSON(fiber.Map{
		"message": "email sent",
	})
	return nil
}
