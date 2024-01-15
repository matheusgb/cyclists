package services

import "github.com/matheusgb/cyclists/src/models/repositories"

type IUser interface {
	// CreateUser(user domains.User) (domains.User, error)
}

type User struct {
	repository repositories.User
}

func Init(service repositories.User) IUser {
	return &User{service}
}
