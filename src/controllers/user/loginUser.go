package controllers

import (
	"github.com/gofiber/fiber/v2"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/user"
	"github.com/matheusgb/cyclists/src/controllers/validators"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/views"
)

func (user *User) LoginUser(ctx *fiber.Ctx) error {
	var request requests.LoginUser

	if err := ctx.BodyParser(&request); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	errValidator := validators.LoginUser(request)
	if errValidator != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": errValidator,
		})
		return nil
	}

	domain := domains.InitLogin(request.Email, request.Password)

	token, err := user.service.LoginUser(*domain)
	if err != nil {
		ctx.Status(401).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	view := views.ConvertStringTokenToResponse(token)
	ctx.Status(200).JSON(view)
	return nil
}
