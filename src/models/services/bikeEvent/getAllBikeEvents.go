package services

import (
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (bikeEvent *BikeEvent) GetAllBikeEvents() ([]entities.BikeEvent, error) {
	return bikeEvent.repository.GetAllBikeEvents()
}
