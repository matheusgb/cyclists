package requests

import "time"

type BikeEvent struct {
	Name                  string    `json:"name"`
	StartPlace            string    `json:"start_place"`
	StartDate             time.Time `json:"start_date"`
	StartDateRegistration time.Time `json:"start_date_registration"`
	EndDateRegistration   time.Time `json:"end_date_registration"`
	Organizer             uint      `json:"organizer"`
	ParticipantsLimit     *uint     `json:"participants_limit"`
	AditionalInformation  *string   `json:"aditional_information"`
}
