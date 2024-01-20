package controllers

import (
	"github.com/gofiber/fiber/v2"
	services "github.com/matheusgb/cyclists/src/models/services/bikeEvent"
)

type BikeEvent struct {
	service services.IBikeEvent
}

type IBikeEvent interface {
	CreateBikeEvent(ctx *fiber.Ctx) error
	UpdateBikeEvent(ctx *fiber.Ctx) error
	// DeleteBikeEvent(ctx *fiber.Ctx) error
	// GetBikeEvent(ctx *fiber.Ctx) error
	// GetAllBikeEvents(ctx *fiber.Ctx) error
}

func Init(service services.IBikeEvent) IBikeEvent {
	return &BikeEvent{service}
}
