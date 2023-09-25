package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://general:a7mod@atlascluster.lmdzbps.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT

var collection *mongo.Collection

// connect eith mongoDB

func init() { // runs only in the beginning, only one time
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success mate")

	collection = client.Database(dbName).Collection(colName)

	// collection instance / reference

	fmt.Println("Collection Instance is ready")

}
