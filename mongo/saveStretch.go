package mongo

import (
	"context"
	"fmt"
	"i9-esa/datatypes"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func nameInCurrentSt(name string, currents map[string][]datatypes.Stretch) primitive.ObjectID {
	for _, exer := range currents["Dynamic"] {
		if exer.Name == name {
			return exer.ID
		}
	}
	for _, exer := range currents["Static"] {
		if exer.Name == name {
			return exer.ID
		}
	}
	return primitive.NilObjectID
}

func SaveStretch(collection *mongo.Collection, stretches []datatypes.Stretch, currents map[string][]datatypes.Stretch) map[string]string {

	ret := map[string]string{}
	for _, stretch := range stretches {

		existingID := nameInCurrentSt(stretch.Name, currents)
		if existingID != primitive.NilObjectID {
			filter := bson.M{"_id": existingID}
			update := bson.M{"$set": stretch}
			updateOptions := options.Update().SetUpsert(true)

			_, err := collection.UpdateOne(context.TODO(), filter, update, updateOptions)
			if err != nil {
				fmt.Println(err)
				continue
			}
			ret[existingID.Hex()] = stretch.Name
		} else {
			insertResult, err := collection.InsertOne(context.Background(), stretch)
			if err != nil {
				fmt.Println(err)
			}
			id := insertResult.InsertedID.(primitive.ObjectID).Hex()
			ret[id] = stretch.Name
		}

	}
	return ret

}
