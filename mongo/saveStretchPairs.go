package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveStretchPairs(collection *mongo.Collection, idMap map[string]string, updates map[string][]string) error {

	for title, list := range updates {
		actualID := nameToID(idMap, title)
		if actualID == "" {
			return errors.New("mismatch with allowed id: title = " + title)
		}

		current := []string{}

		for _, otherName := range list {
			otherID := nameToID(idMap, otherName)
			if otherID == "" {
				return errors.New("mismatch with allowed id: title = " + title + "listed title = " + otherName)
			}

			current = append(current, otherID)
		}

		if err := updateDynamicPairs(collection, actualID, current); err != nil {
			return err
		}

	}

	return nil
}

func nameToID(idMap map[string]string, matchName string) string {
	for id, name := range idMap {
		if name == matchName {
			return id
		}
	}

	return ""
}

func updateDynamicPairs(collection *mongo.Collection, id string, list []string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ObjectID: %v", err)
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"dynamicpairs": list}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update document: %v", err)
	}

	return nil
}
