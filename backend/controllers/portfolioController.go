package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"portfolio/config"
	"portfolio/models"
)

// GetPortfolio retrieves all portfolio entries from the database
func GetPortfolio(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, name, biography, films, awards FROM portfolios")
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve portfolios"})
		return
	}
	defer rows.Close()

	var portfolios []models.Portfolio
	for rows.Next() {
		var portfolio models.Portfolio
		if err := rows.Scan(&portfolio.ID, &portfolio.Name, &portfolio.Biography, &portfolio.Films, &portfolio.Awards); err != nil {
			log.Fatalf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse portfolio"})
			return
		}
		portfolios = append(portfolios, portfolio)
	}

	c.JSON(http.StatusOK, portfolios)
}

// CreatePortfolio adds a new portfolio to the database
func CreatePortfolio(c *gin.Context) {
	var portfolio models.Portfolio
	if err := c.ShouldBindJSON(&portfolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := config.DB.Exec("INSERT INTO portfolios (name, biography, films, awards) VALUES (?, ?, ?, ?)",
		portfolio.Name, portfolio.Biography, portfolio.Films, portfolio.Awards)
	if err != nil {
		log.Fatalf("Error inserting portfolio: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create portfolio"})
		return
	}

	id, _ := result.LastInsertId()
	portfolio.ID = int(id)

	c.JSON(http.StatusOK, portfolio)
}
