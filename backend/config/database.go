package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func config() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	moviesDatabase := client.Database("movies")
	namesCollection := moviesDatabase.Collection("names")

	// Utiliser bson.D pour créer un document BSON
	namesOfMovies, err := namesCollection.InsertOne(ctx, bson.D{
		{Key: "name", Value: "jkghfv"},
		{Key: "title", Value: "lzkesghvjn"},
	})

	if err != nil {
		log.Println(err)
	}

	fmt.Println(namesOfMovies.InsertedID)
}
