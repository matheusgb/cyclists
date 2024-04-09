package services

import (
	domainsP "github.com/matheusgb/cyclists/src/models/domains/pagination"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/user"
)

type IUser interface {
	CreateUser(user domains.User) (entities.User, error)
	UpdateUser(user domains.User) (entities.User, error)
	GetUser(user domains.User) (entities.User, error)
	DeleteUser(user domains.User) (entities.User, error)
	GetAllUsers(pagination *domainsP.Pagination) (*domainsP.Pagination, error)

	LoginUser(user domains.User) (string, entities.User, error)
}

type User struct {
	repository repositories.IUser
}

func Init(service repositories.IUser) IUser {
	return &User{service}
}
