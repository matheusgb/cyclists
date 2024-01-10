package models

import "gorm.io/gorm"

type BikeEvent struct {
	gorm.Model
	Name                  string `gorm:"not null"`
	StartDate             string `gorm:"not null"`
	StartDateRegistration string `gorm:"not null"`
	EndDateRegistration   string `gorm:"not null"`
	AditionalInformation  *string
	StartPlace            string `gorm:"not null"`
	ParticipantsLimit     *uint
	Organizer             uint `gorm:"not null"`
	User                  User `gorm:"foreignKey:Organizer"`
}
