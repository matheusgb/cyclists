package controllers

import (
	"github.com/gofiber/fiber/v2"
	views "github.com/matheusgb/cyclists/src/views/user"
)

func (user *User) GetAllUsers(ctx *fiber.Ctx) error {
	entities, err := user.service.GetAllUsers()
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	view := views.ConvertAllUsersEntityToResponse(entities)
	ctx.Status(200).JSON(view)
	return nil
}
