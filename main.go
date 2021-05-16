package main

import (
	"go-rest-api/database"
	"go-rest-api/routes"
)

func main() {
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	routes.Setup()
}
