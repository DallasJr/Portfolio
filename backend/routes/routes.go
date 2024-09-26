package routes

import (
	"github.com/gin-gonic/gin"
	"portfolio/controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Portfolio Routes
	r.POST("/portfolios", controllers.CreatePortfolio)
	r.GET("/portfolios", controllers.GetPortfolios)
	r.GET("/portfolios/:id", controllers.GetPortfolioByID)
	r.PUT("/portfolios/:id", controllers.UpdatePortfolio)
	r.DELETE("/portfolios/:id", controllers.DeletePortfolio)

	return r
}
