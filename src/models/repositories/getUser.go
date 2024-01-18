package repositories

import (
	"fmt"

	"github.com/matheusgb/cyclists/src/models/domains"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) GetUser(domain domains.User) (entities.User, error) {
	var entity entities.User

	result := user.database.Where("id = ?", domain.ID).First(&entity)
	if result.Error != nil {
		return entity, result.Error
	}
	if result.Row() == nil {
		return entity, fmt.Errorf("user with id %s not found", domain.ID)
	}

	return entity, nil
}
