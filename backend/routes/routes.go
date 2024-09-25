package routes

import (
	"github.com/gin-gonic/gin"
	"portfolio/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Route pour obtenir les informations de base d'Omar Sy
	router.GET("/omar-sy", controllers.GetOmarSyInfo)

	// Route pour obtenir la liste des films
	router.GET("/omar-sy/movies", controllers.GetOmarSyMovies)

	return router
}
