package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"portfolio/config"
	"portfolio/models"
)

var validate = validator.New()

// CreatePortfolio creates a new portfolio
func CreatePortfolio(c *gin.Context) {
	var portfolio models.Portfolio
	if err := c.ShouldBindJSON(&portfolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the input
	if err := validate.Struct(portfolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	config.DB.Create(&portfolio)
	c.JSON(http.StatusOK, portfolio)
}

// GetPortfolios retrieves all portfolios
func GetPortfolios(c *gin.Context) {
	var portfolios []models.Portfolio
	config.DB.Preload("MadeMovies").Preload("PlayedMovies").Find(&portfolios)
	c.JSON(http.StatusOK, portfolios)
}

// GetPortfolioByID retrieves a single portfolio by ID
func GetPortfolioByID(c *gin.Context) {
	id := c.Param("id")
	var portfolio models.Portfolio
	if err := config.DB.Preload("MadeMovies").Preload("PlayedMovies").First(&portfolio, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}
	c.JSON(http.StatusOK, portfolio)
}

// UpdatePortfolio updates a portfolio by ID
func UpdatePortfolio(c *gin.Context) {
	id := c.Param("id")
	var portfolio models.Portfolio
	if err := config.DB.First(&portfolio, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	if err := c.ShouldBindJSON(&portfolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	if err := validate.Struct(portfolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	config.DB.Save(&portfolio)
	c.JSON(http.StatusOK, portfolio)
}

// DeletePortfolio deletes a portfolio by ID
func DeletePortfolio(c *gin.Context) {
	id := c.Param("id")
	var portfolio models.Portfolio
	if err := config.DB.First(&portfolio, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}
	config.DB.Delete(&portfolio)
	c.JSON(http.StatusOK, gin.H{"message": "Portfolio deleted successfully"})
}
