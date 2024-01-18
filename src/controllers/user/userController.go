package controllers

import (
	"github.com/gofiber/fiber/v2"
	services "github.com/matheusgb/cyclists/src/models/services/user"
)

type User struct {
	service services.IUser
}

type IUser interface {
	CreateUser(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
	GetAllUsers(ctx *fiber.Ctx) error
}

func Init(service services.IUser) IUser {
	return &User{service}
}
