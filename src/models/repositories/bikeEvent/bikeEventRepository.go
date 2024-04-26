package repositories

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	domainsP "github.com/matheusgb/cyclists/src/models/domains/pagination"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	"gorm.io/gorm"
)

type IBikeEvent interface {
	CreateBikeEvent(bikeEvent domains.BikeEvent) (entities.BikeEvent, error)
	UpdateBikeEvent(bikeEvent domains.BikeEvent, organizer uint) (entities.BikeEvent, error)
	DeleteBikeEvent(bikeEvent domains.BikeEvent) (entities.BikeEvent, error)
	UpdateBikeEventAdmin(bikeEvent domains.BikeEvent) (entities.BikeEvent, error)
	DeleteBikeEventAdmin(bikeEvent domains.BikeEvent) (entities.BikeEvent, error)
	GetBikeEvent(bikeEvent domains.BikeEvent) (entities.BikeEvent, error)
	GetAllBikeEvents(pagination *domainsP.Pagination, name string) (*domainsP.Pagination, error)
}

type BikeEvent struct {
	database *gorm.DB
}

func Init(database *gorm.DB) IBikeEvent {
	return &BikeEvent{database}
}
