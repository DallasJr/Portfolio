package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Route de test
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Bienvenue dans le backend en Go !"})
	})

	// Démarre le serveur sur le port 8080
	router.Run(":8080")

}
