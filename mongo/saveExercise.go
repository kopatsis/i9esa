package mongo

import (
	"context"
	"fmt"
	"i9-esa/datatypes"

	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func nameInCurrentEx(name string, currents map[string]datatypes.Exercise) primitive.ObjectID {
	for _, exer := range currents {
		if exer.Name == name {
			return exer.ID
		}
	}
	return primitive.NilObjectID
}

func SaveExercise(collection *mongo.Collection, exercises []datatypes.Exercise, currents map[string]datatypes.Exercise) map[string]string {

	ret := map[string]string{}
	for _, exer := range exercises {

		existingID := nameInCurrentEx(exer.Name, currents)
		if existingID != primitive.NilObjectID {
			filter := bson.M{"_id": existingID}
			update := bson.M{"$set": exer}
			fmt.Println(exer.Name, exer.StartQuality)
			updateOptions := options.Update().SetUpsert(true)

			_, err := collection.UpdateOne(context.TODO(), filter, update, updateOptions)
			if err != nil {
				fmt.Println(err)
				continue
			}
			ret[existingID.Hex()] = exer.Name
		} else {
			insertResult, err := collection.InsertOne(context.Background(), exer)
			if err != nil {
				fmt.Println(err)
				continue
			}
			id := insertResult.InsertedID.(primitive.ObjectID).Hex()
			ret[id] = exer.Name
		}

	}

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// cursor, err := collection.Find(ctx, bson.M{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil
	// }
	// defer cursor.Close(ctx)

	// nameToID := map[string]string{}

	// for cursor.Next(ctx) {
	// 	var exer datatypes.RetExercise
	// 	if err := cursor.Decode(&exer); err != nil {
	// 		fmt.Println(err)
	// 		return nil
	// 	}
	// 	nameToID[exer.Name] = exer.ID.Hex()
	// }

	// if err := cursor.Err(); err != nil {
	// 	fmt.Println(err)
	// 	return nil
	// }

	// for name, compatMap := range compats {
	// 	saveID, err := primitive.ObjectIDFromHex(nameToID[name])
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return nil
	// 	}
	// 	saveMap := map[string][2]float32{}

	// 	for compat_name, vals := range compatMap {
	// 		saveMap[nameToID[compat_name]] = vals
	// 	}

	// 	filter := bson.M{"_id": saveID}

	// 	update := bson.M{"$set": bson.M{"compatibles": saveMap}}

	// 	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
	// 		fmt.Println(err)
	// 		return nil
	// 	}
	// }

	return ret

}
