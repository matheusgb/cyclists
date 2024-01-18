package entities

import (
	"time"

	"gorm.io/gorm"
)

type EventUserSubscription struct {
	ID          uint      `gorm:"primarykey"`
	BikeEventID uint      `gorm:"not null"`
	BikeEvent   BikeEvent `gorm:"foreignKey:BikeEventID"`
	UserID      uint      `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
