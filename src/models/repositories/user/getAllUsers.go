package repositories

import (
	domainsP "github.com/matheusgb/cyclists/src/models/domains/pagination"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) GetAllUsers(pag *domainsP.Pagination, email string) (*domainsP.Pagination, error) {
	var entities []entities.User

	result := user.database.Where("LOWER(users.email) LIKE LOWER(?)", "%"+email+"%").Scopes(domainsP.Paginate(&entities, pag, user.database.Where("LOWER(users.email) LIKE LOWER(?)", "%"+email+"%"))).Select("id, name, email, created_at, updated_at").Preload("BikeEvents").Find(&entities)
	if result.Error != nil {
		return pag, result.Error
	}
	pag.Rows = entities

	return pag, nil
}
