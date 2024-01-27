package controllers

import (
	"github.com/gofiber/fiber/v2"
	services "github.com/matheusgb/cyclists/src/models/services/userSubscription"
)

type UserSubscription struct {
	service services.IUserSubscription
}

type IUserSubscription interface {
	CreateUserSubscription(ctx *fiber.Ctx) error
	DeleteUserSubscription(ctx *fiber.Ctx) error
	ListUsersByBikeEvent(ctx *fiber.Ctx) error
}

func Init(service services.IUserSubscription) IUserSubscription {
	return &UserSubscription{service}
}
