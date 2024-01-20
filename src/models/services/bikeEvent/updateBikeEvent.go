package services

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (bikeEvent *BikeEvent) UpdateBikeEvent(domain domains.BikeEvent) (entities.BikeEvent, error) {
	return bikeEvent.repository.UpdateBikeEvent(domain)
}
