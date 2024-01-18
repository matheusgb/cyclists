package services

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) GetUser(domain domains.User) (entities.User, error) {
	entity, err := user.repository.GetUser(domain)
	if err != nil {
		return entity, err
	}
	return entity, nil
}
