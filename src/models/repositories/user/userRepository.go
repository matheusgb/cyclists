package repositories

import (
	domainsP "github.com/matheusgb/cyclists/src/models/domains/pagination"
	domains "github.com/matheusgb/cyclists/src/models/domains/user"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	"gorm.io/gorm"
)

type IUser interface {
	CreateUser(user domains.User) (entities.User, error)
	UpdateUser(user domains.User) (entities.User, error)
	GetUser(user domains.User) (entities.User, error)
	DeleteUser(user domains.User) (entities.User, error)
	GetAllUsers(pag *domainsP.Pagination, email string) (*domainsP.Pagination, error)

	FindUserByEmail(user domains.User) (entities.User, error)
	FindUserByEmailAndPassword(user domains.User) (entities.User, error)
}

type User struct {
	database *gorm.DB
}

func Init(database *gorm.DB) IUser {
	return &User{database}
}
