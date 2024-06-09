package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/user"
	"github.com/matheusgb/cyclists/src/controllers/validators"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
)

func (user *User) CreateUser(ctx *fiber.Ctx) error {
	var request requests.CreateUser

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

	domain := domains.InitCreate(request.Email, request.Password, request.Name)

	entity, err := user.service.CreateUser(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	ctx.Status(201).JSON(fiber.Map{
		"id": fmt.Sprint(entity.ID),
	})
	return nil
}
