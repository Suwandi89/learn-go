package main

import (
	"gin-gorm/database"
	"gin-gorm/routers"
)

func main() {
	database.StartDB()
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
