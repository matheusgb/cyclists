package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheusgb/cyclists/src/models/domains"
)

func (user *User) DeleteUser(ctx *fiber.Ctx) error {
	UserID := ctx.Params("id", "")

	domain := domains.InitID(UserID)

	_, err := user.service.DeleteUser(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	ctx.Status(204)
	return nil
}
