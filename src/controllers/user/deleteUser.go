package controllers

import (
	"github.com/gofiber/fiber/v2"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
)

func (user *User) DeleteUser(ctx *fiber.Ctx) error {
	contextUserID := ctx.Locals("user_id").(string)
	contextUserRole := ctx.Locals("user_role").(string)

	UserID := ctx.Params("id", "")
	if UserID != contextUserID && contextUserRole != "admin" {
		ctx.Status(403).JSON(fiber.Map{
			"message": "you don't have permission to delete this user",
		})
		return nil
	}
	if UserID == "" {
		ctx.Status(400).JSON(fiber.Map{
			"message": "user ID is required",
		})
		return nil
	}

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
