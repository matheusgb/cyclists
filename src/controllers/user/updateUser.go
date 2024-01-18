package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheusgb/cyclists/src/controllers/requests"
	"github.com/matheusgb/cyclists/src/models/domains"
)

func (user *User) UpdateUser(ctx *fiber.Ctx) error {
	var request requests.UpdateUser
	UserID := ctx.Params("id", "")

	if err := ctx.BodyParser(&request); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
		return err
	}

	domain := domains.InitUpdate(request.Name, UserID)

	_, err := user.service.UpdateUser(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	ctx.Status(204)
	return nil
}
