package repositories

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) FindUserByEmail(domain domains.User) (entities.User, error) {
	var entity entities.User

	err := user.database.Where("email = ?", domain.Email).First(&entity).Error
	if err != nil {
		return entity, err
	}

	return entity, nil
}
