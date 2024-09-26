package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username string `json:"username" validate:"required,min=2,max=100"`
	Password string `json:"password" validate:"required"`
}

type Portfolio struct {
	gorm.Model
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Country     string `json:"country" validate:"required"`
	Age         int    `json:"age" validate:"gte=1,lte=150"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}

type MadeMovie struct {
	gorm.Model
	Title       string `json:"title" validate:"required,min=2,max=100"`
	ReleaseDate string `json:"release_date" validate:"required"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type PlayedMovie struct {
	gorm.Model
	Title       string `json:"title" validate:"required,min=2,max=100"`
	ReleaseDate string `json:"release_date" validate:"required"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
