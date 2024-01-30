package controllers

import (
	"github.com/gofiber/fiber/v2"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/user"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/validator"
)

func (user *User) UpdateUser(ctx *fiber.Ctx) error {
	var request requests.UpdateUser
	UserID := ctx.Params("id", "")
	if UserID == "" {
		ctx.Status(400).JSON(fiber.Map{
			"message": "User ID is required",
		})
	}

	if err := ctx.BodyParser(&request); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
		return err
	}

	errValidator := validator.UpdateUser(request)
	if errValidator != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": errValidator,
		})
		return nil
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
