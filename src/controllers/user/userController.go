package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheusgb/cyclists/src/models/services"
)

type User struct {
	service services.IUser
}

type IUser interface {
	CreateUser(ctx *fiber.Ctx) error
}

func Init(service services.IUser) IUser {
	return &User{service}
}
