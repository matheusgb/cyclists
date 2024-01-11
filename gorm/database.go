package gorm

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

func Database() *GormDatabase {
	return &GormDatabase{}
}

func (db *GormDatabase) VerifyConnection(config config.Config) *gorm.DB {
	var client *gorm.DB
	var err error

	if db.client == nil {
		client, err = db.Connect(config)
		if err != nil {
			log.Fatal("Connection attempt failed: ", err)
		}
	}

	return client
}

func (db *GormDatabase) Connect(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection attempt failed: ", err)
		return nil, err
	}

	log.Println("Database connected")
	db.client = client
	return client, nil
}
