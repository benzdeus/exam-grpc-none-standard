package storer_test

import (
	"auth-service/pkg/db"
	"auth-service/pkg/storer"
	"log"
	"testing"
)

func TestCheckPermission(t *testing.T) {

	dbPG, tearDown := db.New()

	defer tearDown()

	repo := storer.NewStore(dbPG)

	per, err := repo.CheckPermission("test")

	if err != nil {
		t.Error(err.Error())
	}

	log.Println(per)
}
