package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/matheusgb/cyclists/src/controllers/requests"
	"github.com/matheusgb/cyclists/src/models/domains"
)

func (user *User) CreateUser(ctx *fiber.Ctx) error {
	var request requests.UserCreate

	if err := ctx.BodyParser(&request); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
		return err
	}

	domain := domains.Init(request.Email, request.Password, request.PasswordConfirmation, request.Name)

	service, err := user.service.CreateUser(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return err
	}

	ctx.Status(201).JSON(fiber.Map{
		"message": fmt.Sprintf("User %s created", service.Name),
	})

	return nil
}
