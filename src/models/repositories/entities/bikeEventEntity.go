package entities

import (
	"time"

	"gorm.io/gorm"
)

type BikeEvent struct {
	gorm.Model
	Name                  string    `gorm:"not null"`
	StartDate             time.Time `gorm:"not null"`
	StartDateRegistration time.Time `gorm:"not null"`
	EndDateRegistration   time.Time `gorm:"not null"`
	StartPlace            string    `gorm:"not null"`
	Organizer             uint      `gorm:"not null"`
	User                  User      `gorm:"foreignKey:Organizer"`
	Participants          []User    `gorm:"many2many:user_subscriptions;"`
	AditionalInformation  *string
	ParticipantsLimit     *uint
}
