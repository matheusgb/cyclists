package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/bikeEvent"
	"github.com/matheusgb/cyclists/src/controllers/validators"
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
)

func (bikeEvent *BikeEvent) CreateBikeEvent(ctx *fiber.Ctx) error {
	var request requests.CreateBikeEvent

	if err := ctx.BodyParser(&request); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	errValidator := validators.BikeEvent(request)
	if errValidator != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": errValidator,
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

	if request.Organizer != organizerId && userRole != "admin" {
		ctx.Status(403).JSON(fiber.Map{
			"message": "you don't have permission to create this bike event on behalf of another user",
		})
		return nil
	}

	domain := domains.InitCreate(request.Name, request.StartPlace, request.AditionalInformation, request.StartDate, request.StartDateRegistration, request.EndDateRegistration, request.Organizer, request.ParticipantsLimit)

	_, err = bikeEvent.service.CreateBikeEvent(*domain)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	ctx.Status(201)
	return nil
}
