package model

import "gorm.io/gorm"

type BikeEvents struct {
	gorm.Model
	Name                  string  `gorm:"not null;" json:"name"`
	StartDate             string  `gorm:"not null;" json:"start_date"`
	StartDateRegistration string  `gorm:"not null;" json:"start_date_registration"`
	EndDateRegistration   string  `gorm:"not null;" json:"end_date_registration"`
	AditionalInformation  *string `json:"aditional_information"`
	StartPlace            string  `gorm:"not null;" json:"start_place"`
	ParticipantsLimit     int     `gorm:"not null;" json:"participants_limit"`
}
