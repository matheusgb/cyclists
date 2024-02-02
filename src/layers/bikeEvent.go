package layers

import (
	controllers "github.com/matheusgb/cyclists/src/controllers/bikeEvent"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/bikeEvent"
	services "github.com/matheusgb/cyclists/src/models/services/bikeEvent"
	"gorm.io/gorm"
)

func InitBikeEvent(database *gorm.DB) controllers.IBikeEvent {
	repository := repositories.Init(database)
	service := services.Init(repository)
	controller := controllers.Init(service)
	return controller
}
