package gorm

import (
	"github.com/matheusgb/cyclists/api/models"
	"github.com/matheusgb/cyclists/config"
	"gorm.io/gorm"
)

func RunMigrations(client *gorm.DB, config config.Config) {
	client.AutoMigrate(&models.BikeEvent{}, &models.User{}, &models.EventUserSubscription{})
}
