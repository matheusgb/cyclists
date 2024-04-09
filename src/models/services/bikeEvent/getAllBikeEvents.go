package services

import (
	domainsP "github.com/matheusgb/cyclists/src/models/domains/pagination"
)

func (bikeEvent *BikeEvent) GetAllBikeEvents(pagination *domainsP.Pagination) (*domainsP.Pagination, error) {
	return bikeEvent.repository.GetAllBikeEvents(pagination)
}
