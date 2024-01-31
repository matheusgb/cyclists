package requests

import "time"

type UpdateBikeEvent struct {
	Name                  string    `json:"name" validate:"omitempty,min=3,max=255"`
	StartPlace            string    `json:"start_place" validate:"omitempty,min=2,max=255"`
	StartDate             time.Time `json:"start_date"`
	StartDateRegistration time.Time `json:"start_date_registration"`
	EndDateRegistration   time.Time `json:"end_date_registration"`
	Organizer             uint      `json:"organizer" validate:"omitempty,gt=0"`
	ParticipantsLimit     *uint     `json:"participants_limit" validate:"omitempty,min=2,max=1000"`
	AditionalInformation  *string   `json:"aditional_information" validate:"omitempty,min=3,max=255"`
}
