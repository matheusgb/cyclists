package controllers

import (
	"github.com/gofiber/fiber/v2"
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	"github.com/matheusgb/cyclists/src/views"
)

func (bikeEvent *BikeEvent) GetBikeEvent(ctx *fiber.Ctx) error {
	bikeEventID := ctx.Params("id", "")
	if bikeEventID == "" {
		ctx.Status(400).JSON(fiber.Map{
			"message": "bike event ID is required",
		})
		return nil
	}

	domain := domains.InitID(bikeEventID)

	entity, err := bikeEvent.service.GetBikeEvent(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	view := views.ConvertBikeEventEntityToResponse(&entity)
	ctx.Status(200).JSON(view)
	return nil
}
