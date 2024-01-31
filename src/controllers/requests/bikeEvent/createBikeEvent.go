package requests

import "time"

type CreateBikeEvent struct {
	Name                  string    `json:"name" validate:"required,min=3,max=255"`
	StartPlace            string    `json:"start_place" validate:"required,min=2,max=255"`
	StartDate             time.Time `json:"start_date" validate:"required"`
	StartDateRegistration time.Time `json:"start_date_registration" validate:"required"`
	EndDateRegistration   time.Time `json:"end_date_registration" validate:"required"`
	Organizer             uint      `json:"organizer" validate:"required"`
	ParticipantsLimit     *uint     `json:"participants_limit" validate:"omitempty,min=2,max=1000"`
	AditionalInformation  *string   `json:"aditional_information" validate:"omitempty,min=3,max=255"`
}
