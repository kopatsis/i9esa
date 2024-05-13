package views

import (
	"context"
	"fmt"
	"i9-esa/datatypes"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetStretchesDB(database *mongo.Database) map[string][]datatypes.Stretch {
	collection := database.Collection("stretch")

	filter := bson.D{}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer cursor.Close(context.Background())

	allstretches := map[string][]datatypes.Stretch{}
	dynamics := []datatypes.Stretch{}
	statics := []datatypes.Stretch{}
	for cursor.Next(context.TODO()) {
		var str datatypes.Stretch
		if err := cursor.Decode(&str); err != nil {
			fmt.Println(err)
			return nil
		}
		if str.Status == "Dynamic" {
			dynamics = append(dynamics, str)
		} else {
			statics = append(statics, str)
		}
	}

	allstretches["Dynamic"] = dynamics
	allstretches["Static"] = statics

	return allstretches
}
