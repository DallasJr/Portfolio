package main

import (
	"portfolio/config"
	"portfolio/controllers"
	"portfolio/routes"
)

func main() {
	config.ConnectDatabase()
	controllers.InitializePortfolio(config.DB)
	r := routes.SetupRoutes()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
