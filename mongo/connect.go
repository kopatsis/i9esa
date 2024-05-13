package mongo

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, *mongo.Database) {

	mong := os.Getenv("MONGO")
	if mong == "" {
		log.Fatal("No mongo variable")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(mong).SetServerAPIOptions(serverAPI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Specify the database and collection
	database := client.Database("i9")

	return client, database
}
