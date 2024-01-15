package services

import (
	"github.com/matheusgb/cyclists/src/models/domains"
	"github.com/matheusgb/cyclists/src/models/repositories"
	"github.com/matheusgb/cyclists/src/views"
)

type IUser interface {
	CreateUser(user domains.User) (views.UserResponse, error)
}

type User struct {
	repository repositories.IUser
}

func Init(service repositories.IUser) IUser {
	return &User{service}
}
