package services

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) UpdateUser(domain domains.User) (entities.User, error) {
	entity, err := user.repository.UpdateUser(domain)
	if err != nil {
		return entity, err
	}
	return entity, nil
}
