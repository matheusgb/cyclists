package controllers

import (
	"github.com/gofiber/fiber/v2"
	domainsP "github.com/matheusgb/cyclists/src/models/domains/pagination"
	"github.com/matheusgb/cyclists/src/views"
)

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
