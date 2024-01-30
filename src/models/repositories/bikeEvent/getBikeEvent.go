package repositories

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (bikeEvent *BikeEvent) GetBikeEvent(domain domains.BikeEvent) (entities.BikeEvent, error) {
	var entity entities.BikeEvent

	result := bikeEvent.database.Joins("User").Preload("Participants").First(&entity, domain.ID)
	if result.Error != nil {
		return entity, result.Error
	}

	return entity, nil
}
