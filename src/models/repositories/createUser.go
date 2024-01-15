package repositories

import (
	"github.com/matheusgb/cyclists/src/models/domains"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	"github.com/matheusgb/cyclists/src/views"
)

func (user *User) CreateUser(domain domains.User) (views.UserResponse, error) {
	var entity entities.User

	err := user.database.Create(&entities.User{
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
		Deleted:  false,
	}).Scan(&entity).Error

	if err != nil {
		return views.UserResponse{}, err
	}

	response := views.ConvertUserEntityToResponse(&entity)

	return *response, nil
}
