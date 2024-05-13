package deletes

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteSretchByID(database *mongo.Database, id string) bool {

	collection := database.Collection("stretch")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}

	filter := bson.M{"_id": objID}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return false
	}

	if result.DeletedCount == 0 {
		return false
	} else {
		return true
	}
}
