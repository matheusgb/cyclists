package repositories

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (bikeEvent *BikeEvent) DeleteBikeEvent(domain domains.BikeEvent) (entities.BikeEvent, error) {
	var entity entities.BikeEvent

	result := bikeEvent.database.Where("id = ? AND organizer = ?", domain.ID, domain.Organizer).Delete(&entity)

	if result.Error != nil {
		return entity, result.Error
	}
	if result.RowsAffected == 0 {
		return entity, fmt.Errorf("bike event with id %s not found", domain.ID)
	} else {
		resultSubs := bikeEvent.database.Where("bike_event_id = ?", domain.ID).Delete(&entities.UserSubscription{})
		if resultSubs.Error != nil {
			return entity, resultSubs.Error
		}
	}

	return entity, nil
}

func (bikeEvent *BikeEvent) DeleteBikeEventAdmin(domain domains.BikeEvent) (entities.BikeEvent, error) {
	var entity entities.BikeEvent

	result := bikeEvent.database.Where("id = ?", domain.ID).Delete(&entity)

	if result.Error != nil {
		return entity, result.Error
	}
	if result.RowsAffected == 0 {
		return entity, fmt.Errorf("bike event with id %s not found", domain.ID)
	} else {
		resultSubs := bikeEvent.database.Where("bike_event_id = ?", domain.ID).Delete(&entities.UserSubscription{})
		if resultSubs.Error != nil {
			return entity, resultSubs.Error
		}
	}

	return entity, nil
}
