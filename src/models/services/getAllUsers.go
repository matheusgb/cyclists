package services

import (
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) GetAllUsers() ([]entities.User, error) {
	return user.repository.GetAllUsers()
}
