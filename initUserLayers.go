package main

import (
	controllers "github.com/matheusgb/cyclists/src/controllers/user"
	repositories "github.com/matheusgb/cyclists/src/models/repositories/user"
	services "github.com/matheusgb/cyclists/src/models/services/user"
	"gorm.io/gorm"
)

func InitUserLayers(database *gorm.DB) controllers.IUser {
	repository := repositories.Init(database)
	service := services.Init(repository)
	controller := controllers.Init(service)
	return controller
}
