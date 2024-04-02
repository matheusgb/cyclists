package services

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (bikeEvent *BikeEvent) DeleteBikeEvent(domain domains.BikeEvent, role string) (entities.BikeEvent, error) {
	if role == "admin" {
		return bikeEvent.repository.DeleteBikeEventAdmin(domain)
	}
	return bikeEvent.repository.DeleteBikeEvent(domain)
}
