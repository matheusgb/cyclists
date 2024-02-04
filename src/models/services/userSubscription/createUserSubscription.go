package services

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (userSubscription *UserSubscription) CreateUserSubscription(domain domains.UserSubscription) (entities.UserSubscription, error) {
	err := userSubscription.repository.FindByUserSubscription(domain)
	if err != nil {
		return entities.UserSubscription{}, err
	}

	err = userSubscription.repository.CheckEventIsFull(domain)
	if err != nil {
		return entities.UserSubscription{}, err
	}

	return userSubscription.repository.CreateUserSubscription(domain)
}
