package repositories

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (userSubscription *UserSubscription) ListUsersByBikeEvent(domain domains.UserSubscription) ([]entities.UserSubscription, error) {
	var userSubscriptions []entities.UserSubscription
	result := userSubscription.database.Joins("User").Where("bike_event_id = ?", domain.BikeEventID).Find(&userSubscriptions)
	if result.Error != nil {
		return nil, fmt.Errorf("error listing users by bike event: %v", result.Error)
	}
	fmt.Println(userSubscriptions)
	return userSubscriptions, nil
}
