package controllers

import (
	"github.com/gofiber/fiber/v2"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/views"
)

// GetUser godoc
//
//	@Summary		Busca um usuário
//	@Description	Retorna um usuário pelo ID
//	@Tags			users
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"ID do usuário"
//	@Success		200	{object}	views.UserResponseWithBikeEvents
//	@Failure		400	{object}	map[string]string
//	@Router			/user/{id} [get]
func (user *User) GetUser(ctx *fiber.Ctx) error {
	UserID := ctx.Params("id", "")
	if UserID == "" {
		ctx.Status(400).JSON(fiber.Map{
			"message": "user ID is required",
		})
		return nil
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
