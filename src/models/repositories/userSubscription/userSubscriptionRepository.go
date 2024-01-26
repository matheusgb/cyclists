package repositories

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	"gorm.io/gorm"
)

type IUserSubscription interface {
	CreateUserSubscription(userSubscription domains.UserSubscription) (entities.UserSubscription, error)
	DeleteUserSubscription(userSubscription domains.UserSubscription) (entities.UserSubscription, error)
}

type UserSubscription struct {
	database *gorm.DB
}

func Init(database *gorm.DB) IUserSubscription {
	return &UserSubscription{database}
}