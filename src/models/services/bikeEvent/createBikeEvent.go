package services

import (
	"fmt"
	"time"

	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (bikeEvent *BikeEvent) CreateBikeEvent(domain domains.BikeEvent) (entities.BikeEvent, error) {
	if domain.StartDate.Before(domain.EndDateRegistration) {
		return entities.BikeEvent{}, fmt.Errorf("the event start date should be after the end date")
	}

	if domain.StartDate.Before(time.Now()) || domain.EndDateRegistration.Before(time.Now()) || domain.StartDateRegistration.Before(time.Now()) {
		return entities.BikeEvent{}, fmt.Errorf("the date registration should be in the future")
	}

	if domain.StartDateRegistration.After(domain.EndDateRegistration) || domain.StartDateRegistration.Equal(domain.EndDateRegistration) {
		return entities.BikeEvent{}, fmt.Errorf("the start date registration should be before the end date registration")
	}

	return bikeEvent.repository.CreateBikeEvent(domain)
}
