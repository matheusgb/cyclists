package services

import (
	"github.com/matheusgb/cyclists/src/models/domains"
	"github.com/matheusgb/cyclists/src/views"
)

func (user *User) CreateUser(domain domains.User) (views.UserResponse, error) {
	return user.repository.CreateUser(domain)
}
