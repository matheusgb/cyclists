package database

import (
	"fmt"
	"log"

	"github.com/matheusgb/cyclists/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDatabase struct {
	client *gorm.DB
}

func CreateGormDatabase() *GormDatabase {
	return &GormDatabase{}
}

func (db *GormDatabase) VerifyConnection(config config.Config) {
	if db.client == nil {
		db.Connect(config)
	} else {
		log.Println("Database already connected")
	}
}

func (db *GormDatabase) Connect(config config.Config) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection attempt failed: ", err)
	}
	db.client = client
	log.Println("Database connected")
}
