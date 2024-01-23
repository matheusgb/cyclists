package controllers

import (
	"github.com/gofiber/fiber/v2"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	views "github.com/matheusgb/cyclists/src/views/user"
)

func (user *User) GetUser(ctx *fiber.Ctx) error {
	UserID := ctx.Params("id", "")
	if UserID == "" {
		ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	domain := domains.InitID(UserID)

	entity, err := user.service.GetUser(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	view := views.ConvertUserEntityToResponse(&entity)
	ctx.Status(200).JSON(view)
	return nil
}
