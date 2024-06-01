package services

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/userSubscription"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/userSubscription"
)

type IUserSubscription interface {
	CreateUserSubscription(userSubscription domains.UserSubscription) (entities.UserSubscription, error)
	DeleteUserSubscription(userSubscription domains.UserSubscription) (entities.UserSubscription, error)
}

type UserSubscription struct {
	repository repositories.IUserSubscription
}

func Init(repository repositories.IUserSubscription) IUserSubscription {
	return &UserSubscription{repository}
}
