package services

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (userSubscription *UserSubscription) CreateUserSubscription(domain domains.UserSubscription) (entities.UserSubscription, error) {
	return userSubscription.repository.CreateUserSubscription(domain)
}
