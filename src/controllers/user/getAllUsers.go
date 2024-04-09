package controllers

import (
	"github.com/gofiber/fiber/v2"
	domainsP "github.com/matheusgb/cyclists/src/models/domains/pagination"
	"github.com/matheusgb/cyclists/src/views"
)

func (user *User) GetAllUsers(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit", 0)
	page := ctx.QueryInt("page", 0)
	sort := ctx.Query("sort", "")

	pagination := &domainsP.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

	response, err := user.service.GetAllUsers(pagination)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	view := views.ConvertAllUsersEntityToResponse(response)
	ctx.Status(200).JSON(view)
	return nil
}
