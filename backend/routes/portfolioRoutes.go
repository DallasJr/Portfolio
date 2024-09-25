package routes

import (
	"github.com/gin-gonic/gin"
	"portfolio/controllers"
)

func PortfolioRoutes(r *gin.Engine) {
	r.GET("/portfolio", controllers.GetPortfolio)
	r.POST("/portfolio", controllers.CreatePortfolio)
}
