package models

import "time"

type EventUserSubscription struct {
	ID               uint      `gorm:"primarykey"`
	BikeEventID      uint      `gorm:"not null"`
	BikeEvent        BikeEvent `gorm:"foreignKey:BikeEventID"`
	UserID           uint      `gorm:"not null"`
	User             User      `gorm:"foreignKey:UserID"`
	SubscriptionDate time.Time `gorm:"not null"`
}
