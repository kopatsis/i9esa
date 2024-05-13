package views

import (
	"context"
	"fmt"
	"i9-esa/datatypes"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetExersDB(database *mongo.Database) map[string]datatypes.Exercise {
	collection := database.Collection("exercise")

	filter := bson.D{}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer cursor.Close(context.Background())

	allexer := map[string]datatypes.Exercise{}

	for cursor.Next(context.TODO()) {
		var exer datatypes.Exercise
		if err := cursor.Decode(&exer); err != nil {
			fmt.Println(err)
			return nil
		}
		allexer[exer.ID.Hex()] = exer
	}

	return allexer
}
