package controllers

import (
	"github.com/gofiber/fiber/v2"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/user"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	views "github.com/matheusgb/cyclists/src/views/user"
)

func (user *User) CreateUser(ctx *fiber.Ctx) error {
	var request requests.CreateUser

	if err := ctx.BodyParser(&request); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
		return err
	}

	domain := domains.InitCreate(request.Email, request.Password, request.PasswordConfirmation, request.Name)

	entity, err := user.service.CreateUser(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	view := views.ConvertUserEntityToResponse(&entity)
	ctx.Status(201).JSON(view)
	return nil
}
