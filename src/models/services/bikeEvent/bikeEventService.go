package services

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	domainsP "github.com/matheusgb/cyclists/src/models/domains/pagination"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/bikeEvent"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

type IBikeEvent interface {
	CreateBikeEvent(bikeEvent domains.BikeEvent) (entities.BikeEvent, error)
	UpdateBikeEvent(bikeEvent domains.BikeEvent, role string, organizer uint) (entities.BikeEvent, error)
	DeleteBikeEvent(bikeEvent domains.BikeEvent, role string) (entities.BikeEvent, error)
	GetBikeEvent(bikeEvent domains.BikeEvent) (entities.BikeEvent, error)
	GetAllBikeEvents(pagination *domainsP.Pagination) (*domainsP.Pagination, error)
}

type BikeEvent struct {
	repository repositories.IBikeEvent
}

func Init(service repositories.IBikeEvent) IBikeEvent {
	return &BikeEvent{service}
}
