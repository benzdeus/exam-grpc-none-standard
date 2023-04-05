package main

import (
	"auth-service/pkg/db"
	"auth-service/pkg/routes"
)

func main() {
	// connect DB
	dbPG, tearDown := db.New()
	defer tearDown()

	// init Router
	routes.InitRoutes(dbPG)

}
