package repositories

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (userSubscription *UserSubscription) DeleteUserSubscription(domain domains.UserSubscription) (entities.UserSubscription, error) {
	var entity entities.UserSubscription

	result := userSubscription.database.Unscoped().Where("bike_event_id = ? AND user_id = ?", domain.BikeEventID, domain.UserID).Delete(&entity)
	if result.Error != nil {
		return entity, result.Error
	}
	if result.RowsAffected == 0 {
		return entity, fmt.Errorf("user subscription with id %d and %d not found", domain.BikeEventID, domain.UserID)
	}

	return entity, nil
}
