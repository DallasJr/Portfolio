package models

import (
	"gorm.io/gorm"
)

// Portfolio model
type Portfolio struct {
	gorm.Model
	Name        string `json:"name" validate:"required,min=2"`
	Country     string `json:"country" validate:"required"`
	Age         int    `json:"age" validate:"required,gt=0"`
	Description string `json:"description" validate:"required"`
	Picture     string `json:"picture" validate:"required"`
}
 
// MadeMovie model
type MadeMovie struct {
	gorm.Model
	Title       string `json:"title" validate:"required"`
	ReleaseDate string `json:"release_date" validate:"required"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image" validate:"required"`
}

// PlayedMovie model
type PlayedMovie struct {
	gorm.Model
	Title       string `json:"title" validate:"required"`
	ReleaseDate string `json:"release_date" validate:"required"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image" validate:"required"`
}
