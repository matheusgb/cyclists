package repositories

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (userSubscription *UserSubscription) CheckEventIsFull(domain domains.UserSubscription) error {
	var count int64
	var participantsLimit *int64

	err := userSubscription.database.
		Model(&entities.UserSubscription{}).
		Where("bike_event_id = ?", domain.BikeEventID).
		Count(&count).
		Error
	if err != nil {
		return err
	}

	err = userSubscription.database.
		Model(&entities.BikeEvent{}).
		Where("id = ?", domain.BikeEventID).
		Select("participants_limit").
		Scan(&participantsLimit).
		Error
	if err != nil {
		return err
	}

	if participantsLimit != nil {
		if count >= *participantsLimit {
			return fmt.Errorf("event is full")
		}
	}
	return nil
}
