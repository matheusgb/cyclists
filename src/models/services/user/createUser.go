package services

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) CreateUser(domain domains.User) (entities.User, error) {
	userFound, err := user.repository.FindUserByEmail(domain)
	if err == nil {
		return userFound, fmt.Errorf("user with email %s already exists", domain.Email)
	} else if err.Error() != "record not found" {
		return userFound, err
	}

	hash, err := domains.HashPassword(domain.Password)
	if err != nil {
		return userFound, fmt.Errorf("invalid password")
	}
	domain.Password = hash

	return user.repository.CreateUser(domain)
}
