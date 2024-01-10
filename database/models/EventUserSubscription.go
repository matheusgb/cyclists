package models

import "time"

type EventUserSubscription struct {
	EventID          uint      `gorm:"not null"`
	BikeEvent        BikeEvent `gorm:"foreignKey:ID"`
	UserID           uint      `gorm:"not null"`
	User             User      `gorm:"foreignKey:ID"`
	SubscriptionDate time.Time `gorm:"not null"`
}
