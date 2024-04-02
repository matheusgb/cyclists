package repositories

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (bikeEvent *BikeEvent) UpdateBikeEvent(domain domains.BikeEvent, organizer uint) (entities.BikeEvent, error) {
	var entity entities.BikeEvent

	result := bikeEvent.database.Where("id = ? AND organizer = ?", domain.ID, organizer).Updates(&entities.BikeEvent{
		Name:                  domain.Name,
		StartPlace:            domain.StartPlace,
		AditionalInformation:  domain.AditionalInformation,
		StartDate:             domain.StartDate,
		StartDateRegistration: domain.StartDateRegistration,
		EndDateRegistration:   domain.EndDateRegistration,
		ParticipantsLimit:     domain.ParticipantsLimit,
		Organizer:             domain.Organizer,
	}).Scan(&entity)

	if result.RowsAffected == 0 {
		return entity, fmt.Errorf("bike event with id %s not found", domain.ID)
	}

	if result.Error != nil {
		return entity, result.Error
	}

	return entity, nil
}

func (bikeEvent *BikeEvent) UpdateBikeEventAdmin(domain domains.BikeEvent) (entities.BikeEvent, error) {
	var entity entities.BikeEvent

	result := bikeEvent.database.Where("id = ?", domain.ID).Updates(&entities.BikeEvent{
		Name:                  domain.Name,
		StartPlace:            domain.StartPlace,
		AditionalInformation:  domain.AditionalInformation,
		StartDate:             domain.StartDate,
		StartDateRegistration: domain.StartDateRegistration,
		EndDateRegistration:   domain.EndDateRegistration,
		ParticipantsLimit:     domain.ParticipantsLimit,
		Organizer:             domain.Organizer,
	}).Scan(&entity)

	if result.RowsAffected == 0 {
		return entity, fmt.Errorf("bike event with id %s not found", domain.ID)
	}

	if result.Error != nil {
		return entity, result.Error
	}

	return entity, nil
}
