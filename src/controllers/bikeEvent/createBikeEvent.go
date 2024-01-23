package controllers

import (
	"github.com/gofiber/fiber/v2"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/bikeEvent"
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	views "github.com/matheusgb/cyclists/src/views/bikeEvent"
)

func (bikeEvent *BikeEvent) CreateBikeEvent(ctx *fiber.Ctx) error {
	var request requests.CreateBikeEvent

	if err := ctx.BodyParser(&request); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
		return err
	}

	domain := domains.InitCreate(request.Name, request.StartPlace, request.AditionalInformation, request.StartDate, request.StartDateRegistration, request.EndDateRegistration, request.Organizer, request.ParticipantsLimit)

	entity, err := bikeEvent.service.CreateBikeEvent(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	view := views.ConvertBikeEventEntityToResponse(&entity)
	ctx.Status(201).JSON(view)
	return nil
}
