package db_test

import (
	"auth-service/pkg/db"
	"testing"
)

func TestConnectDB(t *testing.T) {
	_, tearDown := db.New()
	defer tearDown()
}
