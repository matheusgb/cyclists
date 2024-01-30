package repositories

import (
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (bikeEvent *BikeEvent) GetAllBikeEvents() ([]entities.BikeEvent, error) {
	var entities []entities.BikeEvent

	result := bikeEvent.database.Joins("User").Preload("Participants").Find(&entities)
	if result.Error != nil {
		return entities, result.Error
	}

	return entities, nil
}
