package routes

import (
	"portfolio/config"
	"portfolio/controllers"
	"portfolio/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Configurer les paramètres CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Origine autorisée
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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

	admin.GET("/panel", controllers.GetPanel)

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
