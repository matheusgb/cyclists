package repositories

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (bikeEvent *BikeEvent) CreateBikeEvent(domain domains.BikeEvent) (entities.BikeEvent, error) {
	var entity entities.BikeEvent

	err := bikeEvent.database.Create(&entities.BikeEvent{
		Name:                  domain.Name,
		StartDate:             domain.StartDate,
		Organizer:             domain.Organizer,
		StartDateRegistration: domain.StartDateRegistration,
		EndDateRegistration:   domain.EndDateRegistration,
		StartPlace:            domain.StartPlace,
		ParticipantsLimit:     domain.ParticipantsLimit,
		AditionalInformation:  domain.AditionalInformation,
	}).Scan(&entity).Error

	if err != nil {
		return entity, err
	}

	return entity, nil
}
