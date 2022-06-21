package main

import (
	"assignment2_/database"
	"assignment2_/routes"
)

func main() {
	database.StartDB()

	var PORT = ":8080"

	routes.StartServer().Run(PORT)
}