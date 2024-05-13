package views

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func PrintAllExers(database *mongo.Database) {
	exers := GetExersDB(database)

	if len(exers) == 0 {
		fmt.Println("No exercises in.")
	} else {
		fmt.Printf("All exercises currently in DB: \n\n")

		for _, exer := range exers {
			PrettyPrint(exer)
			fmt.Println()
		}
	}

}

func PrintAllStretches(database *mongo.Database) {
	stretches := GetStretchesDB(database)

	if len(stretches["Dynamic"]) == 0 && len(stretches["Static"]) == 0 {
		fmt.Println("No stretches in.")
	} else {
		fmt.Printf("All stretches currently in DB: \n\n")

		for _, str := range stretches["Dynamic"] {
			fmt.Println("Dynamic: ")
			PrettyPrint(str)
			fmt.Println()
		}

		for _, str := range stretches["Static"] {
			fmt.Println("Static: ")
			PrettyPrint(str)
			fmt.Println()
		}
	}

}

func PrintMatrix(database *mongo.Database) {
	matrix := GetMatrix(database)

	if matrix.ID == primitive.NilObjectID {
		fmt.Println("No matrix in.")
	} else {
		fmt.Printf("Current exercise type matrix: \n\n")

		PrettyPrint(matrix)
	}
}

func PrintOneExer(id string, database *mongo.Database) {
	exer, err := GetOneExerDB(database, id)

	if err == nil {
		fmt.Printf("Requested exercise: \n\n")

		PrettyPrint(exer)
	} else {
		fmt.Println("Invalid id.")
	}
}

func PrintOneStretch(id string, database *mongo.Database) {
	str, err := GetOneStretchDB(database, id)
	if err == nil {
		fmt.Printf("Requested stretch: \n\n")

		PrettyPrint(str)
	} else {
		fmt.Println("Invalid id.")
	}
}

func PrettyPrint(obj interface{}) {
	bytes, _ := json.MarshalIndent(obj, "\t", "\t")
	fmt.Println(string(bytes))
}
