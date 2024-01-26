package services

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/userSubscription"
)

type IUserSubscription interface {
	CreateUserSubscription(user domains.UserSubscription) (entities.UserSubscription, error)
}

type UserSubscription struct {
	repository repositories.IUserSubscription
}

func Init(service repositories.IUserSubscription) IUserSubscription {
	return &UserSubscription{service}
}
