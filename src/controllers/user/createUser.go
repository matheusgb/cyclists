package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheusgb/cyclists/src/controllers/requests"
	"github.com/matheusgb/cyclists/src/models/domains"
	"github.com/matheusgb/cyclists/src/views"
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

	service, err := user.service.CreateUser(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return err
	}

	view := views.ConvertUserEntityToResponse(&service)
	ctx.Status(201).JSON(view)
	return nil
}
