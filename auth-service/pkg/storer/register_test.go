package storer_test

import (
	"auth-service/pkg/db"
	"auth-service/pkg/entities"
	"auth-service/pkg/storer"
	"testing"
)

func TestRegisterRepo(t *testing.T) {
	dbPG, tearDown := db.New()

	defer tearDown()

	repo := storer.NewStore(dbPG)

	user := &entities.User{
		Username:   "test",
		Password:   "test",
		Permission: "admin",
	}

	err := repo.Register(user)
	if err != nil {
		t.Error(err.Error())
	}
}
