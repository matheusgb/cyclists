package repositories

import (
	"fmt"

	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (userSubscription *UserSubscription) FindByUserSubscription(domain domains.UserSubscription) error {
	err := userSubscription.database.Where(domain).First(&entities.UserSubscription{}).Error
	if err == nil {
		return fmt.Errorf("user already subscribed to this event")
	} else if err.Error() != "record not found" {
		return err
	}
	return nil
}
