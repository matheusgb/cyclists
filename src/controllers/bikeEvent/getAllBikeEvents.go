package controllers

import (
	"github.com/gofiber/fiber/v2"
	domainsP "github.com/matheusgb/cyclists/src/models/domains/pagination"
	"github.com/matheusgb/cyclists/src/views"
)

// GetAllBikeEvents godoc
//
//	@Summary		Lista eventos de ciclismo
//	@Description	Lista eventos com paginação e filtro por nome
//	@Tags			bike-events
//	@Produce		json
//	@Security		BearerAuth
//	@Param			limit	query		int		false	"Limite de itens por página"
//	@Param			page	query		int		false	"Página"
//	@Param			sort	query		string	false	"Campo de ordenação"
//	@Param			name	query		string	false	"Filtro por nome"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		400		{object}	map[string]string
//	@Router			/bike-event [get]
func (bikeEvent *BikeEvent) GetAllBikeEvents(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit", 0)
	page := ctx.QueryInt("page", 0)
	sort := ctx.Query("sort", "")
	name := ctx.Query("name", "")

	pagination := &domainsP.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

	response, err := bikeEvent.service.GetAllBikeEvents(pagination, name)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	view := views.ConvertAllBikeEventsEntityToResponse(response)
	ctx.Status(200).JSON(view)
	return nil
}
