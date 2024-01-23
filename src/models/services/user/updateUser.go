package services

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) UpdateUser(domain domains.User) (entities.User, error) {
	return user.repository.UpdateUser(domain)
}
