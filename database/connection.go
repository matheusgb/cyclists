package gorm

import (
	"fmt"
	"log"

	"github.com/matheusgb/cyclists/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (gormDB *GormDatabase) Connect(config config.Config) {
	if gormDB.client == nil {
		gormDB.InitializeClient(config)
	} else {
		log.Fatal("database connection already established")
	}
}

func (gormDB *GormDatabase) InitializeClient(config config.Config) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to open database client connection")
	}

	gormDB.client = client
	log.Println("database client initialized")
}
