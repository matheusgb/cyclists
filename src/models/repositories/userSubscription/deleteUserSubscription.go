package repositories

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (userSubscription *UserSubscription) DeleteUserSubscription(domain domains.UserSubscription) (entities.UserSubscription, error) {
	var entity entities.UserSubscription

	result := userSubscription.database.Where("id = ?", domain.ID).Delete(&entity)
	if result.Error != nil {
		return entity, result.Error
	}
	if result.RowsAffected == 0 {
		return entity, fmt.Errorf("userSubscription with id %s not found", domain.ID)
	}

	return entity, nil
}
