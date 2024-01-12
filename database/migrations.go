package gorm

import (
	"log"

	"github.com/matheusgb/cyclists/config"
	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (gormDB *GormDatabase) RunMigrations(config config.Config) {
	if gormDB.client == nil {
		log.Fatal("Database client not initialized")
	}
	gormDB.client.AutoMigrate(&entities.BikeEvent{}, &entities.User{}, &entities.EventUserSubscription{})
}
