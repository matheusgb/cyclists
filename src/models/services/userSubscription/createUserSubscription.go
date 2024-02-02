package services

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (userSubscription *UserSubscription) CreateUserSubscription(domain domains.UserSubscription) (entities.UserSubscription, error) {
	subscriptrionFound, err := userSubscription.repository.FindByUserSubscription(domain)
	if err == nil {
		return subscriptrionFound, fmt.Errorf("user already subscribed to this event")
	} else if err.Error() != "record not found" {
		return subscriptrionFound, err
	}

	return userSubscription.repository.CreateUserSubscription(domain)
}
