package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID    primitive.ObjectID `bson:"_id"`
	Title *string            `json:"title"`
	Role  *string            `json:"role"`
	Year  *int               `json:"year"`
}
