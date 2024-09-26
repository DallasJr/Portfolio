package main

import (
	"context"
	"log"
	"net/http"
	"portfolio/routes"
)

func main() {
	// Connect to MongoDB
	routes.ConnectDB()
	defer func() {
		if err := routes.Client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// Register handler for /movies
	http.HandleFunc("/movies", routes.GetMovies)

	// Start the server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
