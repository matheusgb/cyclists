package repositories

import (
	domainsP "github.com/matheusgb/cyclists/src/models/domains/pagination"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (bikeEvent *BikeEvent) GetAllBikeEvents(pagination *domainsP.Pagination, name string) (*domainsP.Pagination, error) {
	var entities []entities.BikeEvent

	result := bikeEvent.database.Where("LOWER(bike_events.name) LIKE LOWER(?)", "%"+name+"%").Scopes(domainsP.Paginate(entities, pagination, bikeEvent.database.Where("LOWER(bike_events.name) LIKE LOWER(?)", "%"+name+"%"))).Joins("User").Preload("Participants").Find(&entities)
	if result.Error != nil {
		return pagination, result.Error
	}
	pagination.Rows = entities

	return pagination, nil
}
