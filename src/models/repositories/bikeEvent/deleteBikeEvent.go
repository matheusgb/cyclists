package repositories

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (bikeEvent *BikeEvent) DeleteBikeEvent(domain domains.BikeEvent) (entities.BikeEvent, error) {
	var entity entities.BikeEvent

	result := bikeEvent.database.Where("id = ?", domain.ID).Delete(&entity)
	if result.Error != nil {
		return entity, result.Error
	}

	if result.RowsAffected == 0 {
		return entity, fmt.Errorf("bike event with id %s not found", domain.ID)
	}

	return entity, nil
}
