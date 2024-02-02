package services

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) CreateUser(domain domains.User) (entities.User, error) {
	userFound, err := user.repository.FindUserByEmail(domain)
	if err == nil {
		return userFound, fmt.Errorf("User already exists")
	}

	return user.repository.CreateUser(domain)
}
