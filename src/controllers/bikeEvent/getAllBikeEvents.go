package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheusgb/cyclists/src/views"
)

func (bikeEvent *BikeEvent) GetAllBikeEvents(ctx *fiber.Ctx) error {
	entities, err := bikeEvent.service.GetAllBikeEvents()
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	view := views.ConvertAllBikeEventsEntityToResponse(entities)
	ctx.Status(200).JSON(view)
	return nil
}
