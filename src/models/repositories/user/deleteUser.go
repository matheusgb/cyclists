package repositories

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) DeleteUser(domain domains.User) (entities.User, error) {
	var entity entities.User

	result := user.database.Where("id = ?", domain.ID).Delete(&entity)
	if result.RowsAffected == 0 {
		return entity, fmt.Errorf("user with id %s not found", domain.ID)
	}
	if result.Error != nil {
		return entity, result.Error
	}

	return entity, nil
}
