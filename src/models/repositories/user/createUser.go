package repositories

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) CreateUser(domain domains.User) (entities.User, error) {
	var entity entities.User

	err := user.database.Where("email = ?", domain.Email).First(&entity).Error
	if err == nil {
		return entity, fmt.Errorf("user with email %s already exists", domain.Email)
	}

	err = user.database.Create(&entities.User{
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
	}).Scan(&entity).Error

	if err != nil {
		return entity, err
	}

	return entity, nil
}
