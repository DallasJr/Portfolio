package main

import (
	"log"
	"net/http"
	"portfolio/routes"
)

func main() {
	// Démarrer le routeur et lier les routes
	router := routes.SetupRouter()

	// Démarrer le serveur sur le port 8080
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
