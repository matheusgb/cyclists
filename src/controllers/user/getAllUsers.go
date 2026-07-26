package controllers

import (
	"github.com/gofiber/fiber/v2"
	domainsP "github.com/matheusgb/cyclists/src/models/domains/pagination"
	"github.com/matheusgb/cyclists/src/views"
)

// GetAllUsers godoc
//
//	@Summary		Lista usuários
//	@Description	Lista usuários com paginação e filtro por e-mail
//	@Tags			users
//	@Produce		json
//	@Security		BearerAuth
//	@Param			limit	query		int		false	"Limite de itens por página"
//	@Param			page	query		int		false	"Página"
//	@Param			sort	query		string	false	"Campo de ordenação"
//	@Param			email	query		string	false	"Filtro por e-mail"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		400		{object}	map[string]string
//	@Router			/user [get]
func (user *User) GetAllUsers(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit", 0)
	page := ctx.QueryInt("page", 0)
	sort := ctx.Query("sort", "")
	email := ctx.Query("email", "")

	pagination := &domainsP.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

	response, err := user.service.GetAllUsers(pagination, email)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	view := views.ConvertAllUsersEntityToResponse(response)
	ctx.Status(200).JSON(view)
	return nil
}
