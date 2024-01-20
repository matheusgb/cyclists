package requests

import "time"

type UpdateBikeEvent struct {
	Name                  string    `json:"name"`
	StartPlace            string    `json:"start_place"`
	StartDate             time.Time `json:"start_date"`
	StartDateRegistration time.Time `json:"start_date_registration"`
	EndDateRegistration   time.Time `json:"end_date_registration"`
	ParticipantsLimit     *uint     `json:"participants_limit"`
	AditionalInformation  *string   `json:"aditional_information"`
}
