package storer_test

import (
	"auth-service/pkg/db"
	"auth-service/pkg/storer"
	"log"
	"testing"
)

func TestLogin(t *testing.T) {
	dbPG, tearDown := db.New()

	defer tearDown()

	repo := storer.NewStore(dbPG)
	user, err := repo.Login("test", "test")
	if err != nil {
		t.Error(err.Error())
	}

	log.Printf("%#+v", user)
}
