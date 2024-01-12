package entities

import (
	"database/sql"
	"time"
)

type EventUserSubscription struct {
	ID               uint      `gorm:"primarykey"`
	BikeEventID      uint      `gorm:"not null"`
	BikeEvent        BikeEvent `gorm:"foreignKey:BikeEventID"`
	UserID           uint      `gorm:"not null"`
	User             User      `gorm:"foreignKey:UserID"`
	SubscriptionDate time.Time `gorm:"not null"`
	Deleted          bool      `gorm:"not null"`
	CreatedAt        time.Time
	DeletedAt        sql.NullTime `gorm:"index"`
}
