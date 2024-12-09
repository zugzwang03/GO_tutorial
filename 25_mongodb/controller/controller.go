package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://golang:golang@cluster0.y1s8j.mongodb.net/"
const dbName = "netflix"
const collectionName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongoDB

func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal()
	}
	fmt.Println("MongoDB connection success")

	collection := client.Database(dbName).Collection(collectionName)

	// collection instance
	fmt.Println("Collection instance is ready")
}
