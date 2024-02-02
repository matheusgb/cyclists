package gorm

import (
	"gorm.io/gorm"
)

func (gormDB *GormDatabase) GetClient() *gorm.DB {
	return gormDB.client
}
