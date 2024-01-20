package controllers

import (
	"github.com/gofiber/fiber/v2"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/bikeEvent"
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	views "github.com/matheusgb/cyclists/src/views/bikeEvent"
)

func (bikeEvent *BikeEvent) UpdateBikeEvent(ctx *fiber.Ctx) error {
	var request requests.UpdateBikeEvent

	bikeEventID := ctx.Params("id", "")

	if err := ctx.BodyParser(&request); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
		return err
	}

	domain := domains.InitUpdate(bikeEventID, request.Name, request.StartPlace, request.AditionalInformation, request.StartDate, request.StartDateRegistration, request.EndDateRegistration, request.ParticipantsLimit)

	entity, err := bikeEvent.service.UpdateBikeEvent(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	view := views.ConvertUserEntityToResponse(&entity)
	ctx.Status(201).JSON(view)
	return nil
}
