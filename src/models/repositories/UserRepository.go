package repositories

import (
	"github.com/matheusgb/cyclists/src/models/domains"
	"github.com/matheusgb/cyclists/src/views"
	"gorm.io/gorm"
)

type IUser interface {
	CreateUser(user domains.User) (views.UserResponse, error)
}

type User struct {
	database *gorm.DB
}

func Init(database *gorm.DB) IUser {
	return &User{database}
}
