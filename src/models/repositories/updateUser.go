package repositories

import (
	"github.com/matheusgb/cyclists/src/models/domains"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) UpdateUser(domain domains.User) (entities.User, error) {
	var entity entities.User

	err := user.database.Model(&entities.User{}).Where("id = ?", domain.ID).Updates(entities.User{
		Name: domain.Name,
	}).Scan(&entity).Error

	if err != nil {
		return entity, err
	}

	return entity, nil
}
