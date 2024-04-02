package services

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (bikeEvent *BikeEvent) UpdateBikeEvent(domain domains.BikeEvent, role string, organizer uint) (entities.BikeEvent, error) {
	if role == "admin" {
		return bikeEvent.repository.UpdateBikeEventAdmin(domain)
	}
	return bikeEvent.repository.UpdateBikeEvent(domain, organizer)
}
