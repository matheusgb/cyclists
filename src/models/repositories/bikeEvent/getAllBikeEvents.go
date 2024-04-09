package repositories

import (
	domainsP "github.com/matheusgb/cyclists/src/models/domains/pagination"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (bikeEvent *BikeEvent) GetAllBikeEvents(pagination *domainsP.Pagination) (*domainsP.Pagination, error) {
	var entities []entities.BikeEvent

	result := bikeEvent.database.Scopes(domainsP.Paginate(&entities, pagination, bikeEvent.database)).Joins("User").Preload("Participants").Find(&entities)
	if result.Error != nil {
		return pagination, result.Error
	}
	pagination.Rows = entities

	return pagination, nil
}
