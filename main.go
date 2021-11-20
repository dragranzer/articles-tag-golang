package main

import (
	"my-module/config"
	"my-module/migrate"
	"my-module/routes"
)

func main() {
	//initiateDB
	config.InitDB()
	migrate.AutoMigrate()

	// //initRoutes
	e := routes.New()

	//starting the server
	e.Start(":1234")
}
