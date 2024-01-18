package services

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) DeleteUser(domain domains.User) (entities.User, error) {
	entity, err := user.repository.DeleteUser(domain)
	if err != nil {
		return entity, err
	}
	return entity, nil
}
