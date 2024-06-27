package actions

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"i9-esa/excel"
	mongodb "i9-esa/mongo"
	"i9-esa/views"

	"go.mongodb.org/mongo-driver/mongo"
)

func RunESA(database *mongo.Database) string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Exercise (E), Stretch (S), Both (B), or Matrix(M): ")
	entry, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return ""
	}

	entry = strings.TrimSpace(entry)
	if entry != "S" && entry != "E" && entry != "B" && entry != "M" {
		fmt.Println("Invalid Entry. Out.")
		return ""
	}

	if entry == "M" {
		AddMatrixToDBFull(database)
		return entry
	}

	if entry == "S" || entry == "B" {
		allSts, nameMap := excel.EnterSt()

		collection := database.Collection("stretch")

		currentStretches := views.GetStretchesDB(database)

		insertStretchResults := mongodb.SaveStretch(collection, allSts, currentStretches)

		if err := mongodb.SaveStretchPairs(collection, insertStretchResults, nameMap); err != nil {
			fmt.Println("URGENT: ERROR -- " + err.Error())
		}

		excel.AddStrToXL(insertStretchResults)
	}

	if entry == "E" || entry == "B" {
		allExs := excel.EnterEx()

		collection := database.Collection("exercise")

		currentExers := views.GetExersDB(database)

		insertExerciseResults := mongodb.SaveExercise(collection, allExs, currentExers)

		excel.AddExerToXL(insertExerciseResults)
	}

	return entry

}
