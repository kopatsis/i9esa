package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"strconv"

	"i9-esa/actions"
	"i9-esa/deletes"
	mongodb "i9-esa/mongo"
	"i9-esa/security"
	"i9-esa/views"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	security.GetPassword()

	client, database := mongodb.ConnectDB()
	defer client.Disconnect(context.Background())

	displayMenu()
	userInput := getUserInput()
	for userInput != 0 {
		switch userInput {
		case -1:
			fmt.Println("Invalid entry. Retry.")
		case 1:
			result := actions.RunESA(database)
			if result == "M" {
				views.PrintMatrix(database)
			} else if result == "S" {
				views.PrintAllStretches(database)
			} else if result == "E" {
				views.PrintAllExers(database)
			} else if result == "B" {
				views.PrintAllStretches(database)
				views.PrintAllExers(database)
			} else {
				fmt.Println("Invalid.")
			}
		case 2:
			id := getID()
			deletes.DeleteExerByID(database, id)
			views.PrintAllExers(database)
		case 3:
			id := getID()
			deletes.DeleteSretchByID(database, id)
			views.PrintAllStretches(database)
		case 4:
			views.PrintAllExers(database)
		case 5:
			views.PrintAllStretches(database)
		case 6:
			id := getID()
			views.PrintOneExer(id, database)
		case 7:
			id := getID()
			views.PrintOneStretch(id, database)
		case 8:
			views.PrintMatrix(database)
		default:
			fmt.Println("Error in menu logic. Retry.")
		}
		userInput = getUserInput()
	}

}

func displayMenu() {
	fmt.Println("Options for i9 esa:")
	fmt.Println("0. Quit")
	fmt.Println("1. Push Excel files (exercises, stretches, matrix) to Mongo")
	fmt.Println("2. Delete an exercise by ID")
	fmt.Println("3. Delete a stretch by ID")
	fmt.Println("4. View all exercises")
	fmt.Println("5. View all stretches")
	fmt.Println("6. View one exercise")
	fmt.Println("7. View one stretch")
	fmt.Println("8. View type matrix")
}

func getUserInput() int {
	fmt.Print("Enter your choice (0-8): ")
	var userInput string
	fmt.Scanln(&userInput)
	if num, err := strconv.Atoi(userInput); err == nil {
		if num >= 0 && num <= 8 {
			return num
		}
	}
	return -1
}

func getID() string {
	reader := bufio.NewReader(os.Stdin)
	id := ""
	for id == "" {
		fmt.Print("ID: ")
		idTemp, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading id input:", err)
			continue
		}

		id = strings.TrimSpace(idTemp)
		if id == "" {
			fmt.Println("Error reading id input: Can't be blank")
		}
	}
	return id
}
