package layers

import (
	controllers "github.com/matheusgb/cyclists/src/controllers/userSubscription"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/userSubscription"
	services "github.com/matheusgb/cyclists/src/models/services/userSubscription"
	"gorm.io/gorm"
)

func InitUserSubscription(database *gorm.DB) controllers.IUserSubscription {
	repository := repositories.Init(database)
	service := services.Init(repository)
	controller := controllers.Init(service)
	return controller
}
