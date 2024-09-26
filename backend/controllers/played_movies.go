package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portfolio/config"
	"portfolio/models"
)

func GetPlayedMovies(c *gin.Context) {
	var playedMovies []models.PlayedMovie
	config.DB.Find(&playedMovies)
	c.JSON(http.StatusOK, playedMovies)
}

func CreatePlayedMovie(c *gin.Context) {
	var playedMovie models.PlayedMovie
	if err := c.ShouldBindJSON(&playedMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(playedMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&playedMovie)
	c.JSON(http.StatusOK, playedMovie)
}

func UpdatePlayedMovie(c *gin.Context) {
	var playedMovie models.PlayedMovie
	id := c.Param("id")
	if err := config.DB.First(&playedMovie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unknown movie"})
		return
	}
	if err := c.ShouldBindJSON(&playedMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(playedMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&playedMovie)
	c.JSON(http.StatusOK, playedMovie)
}

func DeletePlayedMovie(c *gin.Context) {
	var playedMovie models.PlayedMovie
	id := c.Param("id")
	if err := config.DB.First(&playedMovie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unknown movie"})
		return
	}
	config.DB.Delete(&playedMovie)
	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
