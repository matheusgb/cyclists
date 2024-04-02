package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string      `gorm:"not null"`
	Email      string      `gorm:"not null"`
	Password   string      `gorm:"not null"`
	BikeEvents []BikeEvent `gorm:"many2many:user_subscriptions;"`
	Role       string      `gorm:"default: 'user'"` // user or admin
}
