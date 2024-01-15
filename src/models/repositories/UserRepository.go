package repositories

import (
	"gorm.io/gorm"
)

type IUser interface {
	// CreateUser(user domains.UserDomain) entities.User
}

type User struct {
	database *gorm.DB
}

func InitUserRepository(database *gorm.DB) IUser {
	return &User{database}
}
