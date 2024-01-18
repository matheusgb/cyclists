package views

import (
	"time"

	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

type BikeEventResponse struct {
	ID                    uint      `json:"id"`
	Name                  string    `json:"name"`
	StartPlace            string    `json:"start_place"`
	StartDate             time.Time `json:"start_date"`
	StartDateRegistration time.Time `json:"start_date_registration"`
	EndDateRegistration   time.Time `json:"end_date_registration"`
	Organizer             uint      `json:"organizer"`
	ParticipantsLimit     *uint     `json:"participants_limit"`
	AditionalInformation  *string   `json:"aditional_information"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

func ConvertUserEntityToResponse(bikeEventEntity *entities.BikeEvent) *BikeEventResponse {
	return &BikeEventResponse{
		ID:                    bikeEventEntity.ID,
		Name:                  bikeEventEntity.Name,
		StartPlace:            bikeEventEntity.StartPlace,
		StartDate:             bikeEventEntity.StartDate,
		StartDateRegistration: bikeEventEntity.StartDateRegistration,
		EndDateRegistration:   bikeEventEntity.EndDateRegistration,
		Organizer:             bikeEventEntity.Organizer,
		ParticipantsLimit:     bikeEventEntity.ParticipantsLimit,
		AditionalInformation:  bikeEventEntity.AditionalInformation,
		CreatedAt:             bikeEventEntity.CreatedAt,
		UpdatedAt:             bikeEventEntity.UpdatedAt,
	}
}
