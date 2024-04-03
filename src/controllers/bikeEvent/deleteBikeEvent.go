package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
)

func (bikeEvent *BikeEvent) DeleteBikeEvent(ctx *fiber.Ctx) error {
	bikeEventID := ctx.Params("id", "")
	if bikeEventID == "" {
		ctx.Status(400).JSON(fiber.Map{
			"message": "bike event ID is required",
		})
		return nil
	}

	userRole := ctx.Locals("user_role").(string)
	uintOrganizerId, err := strconv.ParseUint(ctx.Locals("user_id").(string), 10, 64)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"message": "internal error",
		})
	}
	organizerId := uint(uintOrganizerId)

	domain := domains.InitID(bikeEventID)
	domain.Organizer = organizerId

	_, err = bikeEvent.service.DeleteBikeEvent(*domain, userRole)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	ctx.Status(204)
	return nil
}
