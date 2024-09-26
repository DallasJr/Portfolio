package main

import (
	"portfolio/config"
	"portfolio/routes"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRoutes()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
