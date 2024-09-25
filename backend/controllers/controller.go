package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portfolio/models"
)

// Récupérer les informations de base d'Omar Sy
func GetOmarSyInfo(c *gin.Context) {
	omar := models.OmarSy{
		Name:        "Omar Sy",
		Description: "Omar Sy est un acteur français, connu pour son rôle dans 'Intouchables'.",
	}

	c.JSON(http.StatusOK, omar)
}

// Récupérer la liste des films d'Omar Sy
func GetOmarSyMovies(c *gin.Context) {
	movies := models.GetMovies()

	c.JSON(http.StatusOK, movies)
}
