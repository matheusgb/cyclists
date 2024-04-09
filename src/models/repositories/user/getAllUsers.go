package repositories

import (
	domainsP "github.com/matheusgb/cyclists/src/models/domains/pagination"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (user *User) GetAllUsers(pagination *domainsP.Pagination) (*domainsP.Pagination, error) {
	var entities []entities.User

	result := user.database.Scopes(domainsP.Paginate(&entities, pagination, user.database)).Select("id, name, email, created_at, updated_at").Preload("BikeEvents").Find(&entities)
	if result.Error != nil {
		return pagination, result.Error
	}
	pagination.Rows = entities

	return pagination, nil
}
