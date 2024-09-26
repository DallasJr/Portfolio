package main

import (
	"portfolio/config"
	"portfolio/routes"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRoutes()
	r.Run(":8080")
}
