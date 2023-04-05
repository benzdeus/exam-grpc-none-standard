package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, func()) {

	// postgrest
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=auth_service port=5432 sslmode=disable TimeZone=Asia/Bangkok"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db, func() {
		db, err := db.DB()
		if err != nil {
			panic(err)
		}
		db.Close()
	}
}
