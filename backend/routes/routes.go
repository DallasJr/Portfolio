package routes

import (
	"github.com/gin-gonic/gin"
	"portfolio/controllers"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/portfolio", controllers.GetPortfolio)
	router.POST("/portfolio", controllers.CreatePortfolio)
	router.PUT("/portfolio", controllers.UpdatePortfolio)

	router.GET("/made-movies", controllers.GetMadeMovies)
	router.POST("/made-movies", controllers.CreateMadeMovie)
	router.PUT("/made-movies/:id", controllers.UpdateMadeMovie)
	router.DELETE("/made-movies/:id", controllers.DeleteMadeMovie)

	router.GET("/played-movies", controllers.GetPlayedMovies)
	router.POST("/played-movies", controllers.CreatePlayedMovie)
	router.PUT("/played-movies/:id", controllers.UpdatePlayedMovie)
	router.DELETE("/played-movies/:id", controllers.DeletePlayedMovie)
	return router
}
