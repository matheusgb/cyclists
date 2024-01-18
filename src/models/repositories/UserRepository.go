package repositories

import (
	"github.com/matheusgb/cyclists/src/models/domains"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	"gorm.io/gorm"
)

type IUser interface {
	CreateUser(user domains.User) (entities.User, error)
	UpdateUser(user domains.User) (entities.User, error)
	GetUser(user domains.User) (entities.User, error)
	DeleteUser(user domains.User) (entities.User, error)
	GetAllUsers() ([]entities.User, error)
}

type User struct {
	database *gorm.DB
}

func Init(database *gorm.DB) IUser {
	return &User{database}
}
