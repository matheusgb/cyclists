package repositories

import (
	"fmt"

	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (bikeEvent *BikeEvent) GetAllBikeEvents() ([]entities.BikeEvent, error) {
	var entities []entities.BikeEvent

	result := bikeEvent.database.Select("id, name, start_date, start_date_registration, end_date_registration, aditional_information, start_place, participants_limit, organizer, created_at, updated_at").Find(&entities)
	if result.Error != nil {
		return entities, result.Error
	}

	if entities == nil {
		return entities, fmt.Errorf("no bike events found")
	}

	return entities, nil
}
