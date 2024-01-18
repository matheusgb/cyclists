package services

import (
	domains "github.com/matheusgb/cyclists/src/models/domains/bikeEvent"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/bikeEvent"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

type IBikeEvent interface {
	CreateBikeEvent(bikeEvent domains.BikeEvent) (entities.BikeEvent, error)
}

type BikeEvent struct {
	repository repositories.IBikeEvent
}

func Init(service repositories.IBikeEvent) IBikeEvent {
	return &BikeEvent{service}
}
