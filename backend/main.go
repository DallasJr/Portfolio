package main

import (
	"github.com/gin-gonic/gin"
	"portfolio/config"
	"portfolio/routes"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()
	routes.PortfolioRoutes(r)

	r.Run(":8080") // Start server on port 8080
}
