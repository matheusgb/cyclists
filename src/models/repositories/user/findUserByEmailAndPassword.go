package repositories

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) FindUserByEmailAndPassword(domain domains.User) (entities.User, error) {
	var entity entities.User

	err := user.database.Where("email = ? AND password = ?", domain.Email, domain.Password).First(&entity).Error
	if err != nil {
		return entity, err
	}

	return entity, nil
}
