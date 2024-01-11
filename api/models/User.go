package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name                 string `gorm:"not null" json:"name"`
	Email                string `gorm:"not null" json:"email"`
	Password             string `gorm:"not null" json:"password"`
	PasswordConfirmation string `gorm:"-" json:"password_confirmation"`
}

type omit *struct{}

type UserResponse struct {
	*User
	Password omit `json:"password,omitempty"`
}
