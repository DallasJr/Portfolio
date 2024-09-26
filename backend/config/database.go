package config

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"portfolio/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("portfolio.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database!")
	}
	log.Println("Database connected successfully.")
	err = DB.AutoMigrate(&models.Portfolio{}, &models.MadeMovie{}, &models.PlayedMovie{}, &models.Admin{})
	if err != nil {
		return
	}
	fmt.Println("Database connection successful.")
}
