package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
	Deleted  bool   `gorm:"not null"`
}
