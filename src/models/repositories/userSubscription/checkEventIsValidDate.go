package repositories

import (
	"fmt"
	"time"

	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (userSubscription *UserSubscription) CheckEventIsValidDate(domain domains.UserSubscription) error {
	var entity entities.BikeEvent

	result := userSubscription.database.Where(domain.BikeEventID).First(&entity)
	if result.Error != nil {
		return result.Error
	}
	if entity.StartDateRegistration.Before(time.Now()) && entity.EndDateRegistration.After(time.Now()) {
		return nil
	} else {
		return fmt.Errorf("the event subscription is out of date")
	}
}
