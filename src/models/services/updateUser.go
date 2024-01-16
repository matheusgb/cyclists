package services

import (
	"errors"

	"github.com/matheusgb/cyclists/src/models/domains"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) UpdateUser(domain domains.User) (entities.User, error) {
	if domain.ID == "" {
		return entities.User{}, errors.New("ID is required")
	}

	return user.repository.UpdateUser(domain)
}
