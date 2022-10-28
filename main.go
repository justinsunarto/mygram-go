package main

import (
	"assignment_4/database"
	"assignment_4/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
