package routes

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client *mongo.Client

func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Change the URI as needed
	var err error
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check connection
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
}
