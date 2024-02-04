package repositories

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (userSubscription *UserSubscription) CheckEventIsFull(domain domains.UserSubscription) (entities.UserSubscription, error) {
	var entity entities.UserSubscription

	// err := userSubscription.database.Where(domain).First(&entity).Error
	// if err != nil {
	// 	return entity, err
	// }

	return entity, nil
}
