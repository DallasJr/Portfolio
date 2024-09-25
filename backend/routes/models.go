package routes

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"portfolio/models"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	collection := Client.Database("portfolio").Collection("movies")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var movies []models.Movie
	if err = cursor.All(context.TODO(), &movies); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
