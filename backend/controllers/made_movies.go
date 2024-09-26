package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portfolio/config"
	"portfolio/models"
)

func GetMadeMovies(c *gin.Context) {
	var madeMovies []models.MadeMovie
	config.DB.Find(&madeMovies)
	c.JSON(http.StatusOK, madeMovies)
}

func CreateMadeMovie(c *gin.Context) {
	var madeMovie models.MadeMovie
	if err := c.ShouldBindJSON(&madeMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(madeMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&madeMovie)
	c.JSON(http.StatusOK, madeMovie)
}

func UpdateMadeMovie(c *gin.Context) {
	var madeMovie models.MadeMovie
	id := c.Param("id")
	if err := config.DB.First(&madeMovie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unknown movie"})
		return
	}
	if err := c.ShouldBindJSON(&madeMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(madeMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&madeMovie)
	c.JSON(http.StatusOK, madeMovie)
}

func DeleteMadeMovie(c *gin.Context) {
	var madeMovie models.MadeMovie
	id := c.Param("id")
	if err := config.DB.First(&madeMovie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unknown movie"})
		return
	}
	config.DB.Delete(&madeMovie)
	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
