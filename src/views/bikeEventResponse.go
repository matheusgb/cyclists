package views

import (
	"time"

	domainsP "github.com/matheusgb/cyclists/src/models/domains/pagination"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

type BikeEventResponse struct {
	ID                   uint    `json:"id"`
	Name                 string  `json:"name"`
	AditionalInformation *string `json:"aditional_information"`
}

type BikeEventResponseWithParticipantsAndOrganizer struct {
	BikeEventResponse
	StartPlace            string         `json:"start_place"`
	StartDate             time.Time      `json:"start_date"`
	StartDateRegistration time.Time      `json:"start_date_registration"`
	EndDateRegistration   time.Time      `json:"end_date_registration"`
	ParticipantsLimit     *uint          `json:"participants_limit"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	Organizer             UserResponse   `json:"organizer"`
	Participants          []UserResponse `json:"participants"`
}

func ConvertBikeEventEntityToResponseUser(bikeEventEntity *entities.BikeEvent) *BikeEventResponse {
	return &BikeEventResponse{
		ID:                   bikeEventEntity.ID,
		Name:                 bikeEventEntity.Name,
		AditionalInformation: bikeEventEntity.AditionalInformation,
	}
}

func ConvertBikeEventEntityToResponse(bikeEventEntity *entities.BikeEvent) *BikeEventResponseWithParticipantsAndOrganizer {
	return &BikeEventResponseWithParticipantsAndOrganizer{
		BikeEventResponse: BikeEventResponse{
			ID:                   bikeEventEntity.ID,
			Name:                 bikeEventEntity.Name,
			AditionalInformation: bikeEventEntity.AditionalInformation,
		},
		StartPlace:            bikeEventEntity.StartPlace,
		StartDate:             bikeEventEntity.StartDate,
		StartDateRegistration: bikeEventEntity.StartDateRegistration,
		EndDateRegistration:   bikeEventEntity.EndDateRegistration,
		CreatedAt:             bikeEventEntity.CreatedAt,
		UpdatedAt:             bikeEventEntity.UpdatedAt,
		ParticipantsLimit:     bikeEventEntity.ParticipantsLimit,
		Organizer:             *ConvertUserEntityToResponseBikeEvent(&bikeEventEntity.User),
		Participants:          ConvertAllUsersEntityToResponseBikeEvent(bikeEventEntity.Participants),
	}
}

func ConvertAllBikeEventsEntityToResponseUser(bikeEventEntity []entities.BikeEvent) []BikeEventResponse {
	var bikeEventsResponse []BikeEventResponse
	for _, bikeEventEntity := range bikeEventEntity {
		bikeEventsResponse = append(bikeEventsResponse, *ConvertBikeEventEntityToResponseUser(&bikeEventEntity))
	}
	return bikeEventsResponse
}

func ConvertAllBikeEventsEntityToResponse(response *domainsP.Pagination) *domainsP.Pagination {
	var bikeEventsResponse []BikeEventResponseWithParticipantsAndOrganizer
	for _, response := range response.Rows.([]entities.BikeEvent) {
		bikeEventsResponse = append(bikeEventsResponse, *ConvertBikeEventEntityToResponse(&response))
	}

	response.Rows = bikeEventsResponse

	return response
}
