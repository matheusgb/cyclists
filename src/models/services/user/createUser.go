package services

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) CreateUser(domain domains.User) (entities.User, error) {
	userFound, err := user.repository.FindUserByEmail(domain)
	if err == nil {
		return userFound, fmt.Errorf("user already exists")
	} else if err.Error() != "record not found" {
		return userFound, err
	}

	return user.repository.CreateUser(domain)
}
