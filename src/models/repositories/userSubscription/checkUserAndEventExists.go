package repositories

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (userSubscription *UserSubscription) CheckUserAndEventExists(domain domains.UserSubscription) error {
	err := userSubscription.database.Where(domain.UserID).First(&entities.User{}).Error
	if err != nil {
		return fmt.Errorf("user not found")
	}

	err = userSubscription.database.Where(domain.BikeEventID).First(&entities.BikeEvent{}).Error
	if err != nil {
		return fmt.Errorf("event not found")
	}

	return nil
}
