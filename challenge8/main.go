package main

import (
	"challenge8/database"
	"challenge8/routers"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(":8080")
}
