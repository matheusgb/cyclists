package controllers

import (
	"github.com/gofiber/fiber/v2"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/bikeEvent"
	"github.com/matheusgb/cyclists/src/controllers/validators"
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
)

func (bikeEvent *BikeEvent) UpdateBikeEvent(ctx *fiber.Ctx) error {
	var request requests.UpdateBikeEvent

	bikeEventID := ctx.Params("id", "")
	if bikeEventID == "" {
		ctx.Status(400).JSON(fiber.Map{
			"message": "BikeEvent ID is required",
		})
		return nil
	}

	if err := ctx.BodyParser(&request); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	errValidator := validators.BikeEvent[requests.UpdateBikeEvent](request)
	if errValidator != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": errValidator,
		})
		return nil
	}

	domain := domains.InitUpdate(bikeEventID, request.Name, request.StartPlace, request.AditionalInformation, request.StartDate, request.StartDateRegistration, request.EndDateRegistration, request.ParticipantsLimit, request.Organizer)

	_, err := bikeEvent.service.UpdateBikeEvent(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	ctx.Status(204)
	return nil
}
