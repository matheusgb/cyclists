package gorm

import (
	"log"

	"github.com/matheusgb/cyclists/src/models/repositories/entities"
)

func (gormDB *GormDatabase) RunMigrations() {
	if gormDB.client == nil {
		log.Fatal("Database client not initialized")
	}
	gormDB.client.AutoMigrate(&entities.BikeEvent{}, &entities.User{}, &entities.UserSubscription{})
}
