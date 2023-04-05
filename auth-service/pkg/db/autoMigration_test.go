package db_test

import (
	"auth-service/pkg/db"
	"testing"
)

func TestAutoMigration(t *testing.T) {
	dbPG, tearDown := db.New()
	defer tearDown()

	err := db.AutoMigration(dbPG)

	if err != nil {
		t.Error(err.Error())
	}
}
