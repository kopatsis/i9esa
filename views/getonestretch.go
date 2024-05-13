package views

import (
	"context"
	"fmt"
	"i9-esa/datatypes"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetOneStretchDB(database *mongo.Database, str string) (datatypes.Stretch, error) {

	var result datatypes.Stretch

	collection := database.Collection("stretch")

	objID, err := primitive.ObjectIDFromHex(str)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	filter := bson.M{"_id": objID}

	if err := collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		fmt.Println(err)
		return result, err
	}

	return result, nil

}
