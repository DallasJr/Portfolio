package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"portfolio/config"
	"portfolio/models"
)

var validate = validator.New()

const defaultUsername = "admin"
const defaultPassword = "admin"

// Crypter le mot de passe
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Vérifier un mot de passe crypté
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func InitializePortfolio(db *gorm.DB) {
	fmt.Println("Creating admin login")
	var admin models.Admin
	db.First(&admin)
	if admin.Username == "" {
		hashedPassword, _ := HashPassword(defaultPassword)
		admin = models.Admin{
			Username: defaultUsername,
			Password: hashedPassword,
		}
		db.Create(&admin)
	}

	fmt.Println("Creating portfolio")
	var portfolio models.Portfolio
	db.First(&portfolio)
	if portfolio.Name == "" {
		portfolio = models.Portfolio{
			Name:        "Name Surname",
			Country:     "France",
			Age:         20,
			Description: "This is a default portfolio created on startup.",
			Picture:     "default_picture.jpg",
		}
		db.Create(&portfolio)
		fmt.Println("Default portfolio created.")
	}
}

func AdminLogin(c *gin.Context) {
	var input models.Admin
	var admin models.Admin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	db.First(&admin)
	if admin.Username != input.Username || !CheckPasswordHash(input.Password, admin.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect username or password"})
		return
	}
	session := sessions.Default(c)
	session.Set("admin", admin.Username)
	err := session.Save()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func AdminLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func GetPanel(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func GetPortfolio(c *gin.Context) {
	var portfolio models.Portfolio
	if err := config.DB.First(&portfolio).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}
	c.JSON(http.StatusOK, portfolio)
}

func UpdatePortfolio(c *gin.Context) {
	var portfolio models.Portfolio
	if err := config.DB.First(&portfolio).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}
	if err := c.ShouldBindJSON(&portfolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(portfolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&portfolio)
	c.JSON(http.StatusOK, portfolio)
}
