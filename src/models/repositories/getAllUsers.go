package repositories

import (
	"fmt"

	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) GetAllUsers() ([]entities.User, error) {
	var entities []entities.User

	result := user.database.Select("id, name, email, created_at, updated_at").Find(&entities)
	if result.Error != nil {
		return entities, result.Error
	}

	if entities == nil {
		return entities, fmt.Errorf("no users found")
	}

	return entities, nil
}
