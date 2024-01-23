package repositories

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
	"gorm.io/gorm"
)

type IBikeEvent interface {
	CreateBikeEvent(bikeEvent domains.BikeEvent) (entities.BikeEvent, error)
	UpdateBikeEvent(bikeEvent domains.BikeEvent) (entities.BikeEvent, error)
	DeleteBikeEvent(bikeEvent domains.BikeEvent) (entities.BikeEvent, error)
	GetBikeEvent(bikeEvent domains.BikeEvent) (entities.BikeEvent, error)
	GetAllBikeEvents() ([]entities.BikeEvent, error)
}

type BikeEvent struct {
	database *gorm.DB
}

func Init(database *gorm.DB) IBikeEvent {
	return &BikeEvent{database}
}
