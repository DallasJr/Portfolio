package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type OmarSy struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
}

type Movie struct {
	Title string `bson:"title"`
	Role  string `bson:"role"`
	Year  int    `bson:"year"`
}

func GetMovies() []Movie {
	return []Movie{
		{Title: "Intouchables", Role: "Driss", Year: 2011},
		{Title: "Jurassic World", Role: "Barry", Year: 2015},
		{Title: "Lupin", Role: "Assane Diop", Year: 2021},
	}
}
