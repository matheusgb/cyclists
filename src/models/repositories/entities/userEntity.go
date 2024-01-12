package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name                 string `gorm:"not null"`
	Email                string `gorm:"not null"`
	Password             string `gorm:"not null"`
	PasswordConfirmation string `gorm:"-" `
	Deleted              bool   `gorm:"not null" `
}
