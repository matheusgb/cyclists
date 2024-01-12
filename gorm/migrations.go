package gorm

import (
	"log"

	"github.com/matheusgb/cyclists/api/models"
	"github.com/matheusgb/cyclists/config"
)

func (gormDB *GormDatabase) RunMigrations(config config.Config) {
	if gormDB.client == nil {
		log.Fatal("Database client not initialized")
	}
	gormDB.client.AutoMigrate(&models.BikeEvent{}, &models.User{}, &models.EventUserSubscription{})
}
