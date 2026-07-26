package controllers

import (
	"github.com/gofiber/fiber/v2"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/user"
	"github.com/matheusgb/cyclists/src/controllers/validators"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
)

// UpdateUser godoc
//
//	@Summary		Atualiza um usuário
//	@Description	Atualiza os dados de um usuário. Requer o próprio usuário ou privilégio de admin
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path	string				true	"ID do usuário"
//	@Param			request	body	requests.UpdateUser	true	"Dados a atualizar"
//	@Success		204
//	@Failure		400	{object}	map[string]string
//	@Failure		403	{object}	map[string]string
//	@Router			/user/{id} [patch]
func (user *User) UpdateUser(ctx *fiber.Ctx) error {
	var request requests.UpdateUser
	contextUserID := ctx.Locals("user_id").(string)
	contextUserRole := ctx.Locals("user_role").(string)

	UserID := ctx.Params("id", "")
	if UserID != contextUserID && contextUserRole != "admin" {
		ctx.Status(403).JSON(fiber.Map{
			"message": "you don't have permission to update this user",
		})
		return nil
	}
	if UserID == "" {
		ctx.Status(400).JSON(fiber.Map{
			"message": "user ID is required",
		})
		return nil
	}

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

	domain := domains.InitUpdate(request.Name, UserID)

	_, err := user.service.UpdateUser(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	ctx.Status(204)
	return nil
}
