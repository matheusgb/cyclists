package gorm

import (
	"github.com/matheusgb/cyclists/config"
	"gorm.io/gorm"
)

type IDatabase interface {
	Connect(config config.Config)
	InitializeClient(config config.Config)
	RunMigrations(config config.Config)
}

type GormDatabase struct {
	client *gorm.DB
}

func Init() IDatabase {
	return &GormDatabase{}
}
