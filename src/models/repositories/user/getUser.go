package repositories

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) GetUser(domain domains.User) (entities.User, error) {
	var entity entities.User

	result := user.database.Where("id = ?", domain.ID).Preload("BikeEvents").First(&entity)
	if result.Error != nil {
		return entity, result.Error
	}

	return entity, nil
}
