package db

import (
	"auth-service/pkg/entities"

	"gorm.io/gorm"
)

func AutoMigration(db *gorm.DB) error {
	return db.AutoMigrate(&entities.User{})
}
