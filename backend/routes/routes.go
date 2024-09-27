package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"portfolio/config"
	"portfolio/controllers"
	"portfolio/middleware"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.Use(DBMiddleware(config.DB))
	router.POST("/login", controllers.AdminLogin)
	router.POST("/logout", controllers.AdminLogout)

	// Accessible seulement par admin
	admin := router.Group("/admin")
	admin.Use(middleware.AuthRequired())

	admin.PUT("/portfolio", controllers.UpdatePortfolio)

	admin.POST("/made-movies", controllers.CreateMadeMovie)
	admin.PUT("/made-movies/:id", controllers.UpdateMadeMovie)
	admin.DELETE("/made-movies/:id", controllers.DeleteMadeMovie)

	admin.POST("/played-movies", controllers.CreatePlayedMovie)
	admin.PUT("/played-movies/:id", controllers.UpdatePlayedMovie)
	admin.DELETE("/played-movies/:id", controllers.DeletePlayedMovie)
	//////////////////////////////////

	router.GET("/portfolio", controllers.GetPortfolio)
	router.GET("/made-movies", controllers.GetMadeMovies)
	router.GET("/played-movies", controllers.GetPlayedMovies)
	return router
}

func DBMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
