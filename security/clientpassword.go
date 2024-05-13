package security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func verifyPassword(password string) error {

	matches := "$2a$10$X/UzmHX71oAw2F87OrwRaObXuUc2ws8mNYnMWvBuHAjV4TOwPHGOm"

	return bcrypt.CompareHashAndPassword([]byte(matches), []byte(password))
}

func GetPassword() {
	fmt.Printf("Enter i9 ESA password: ")

	var userInput string
	fmt.Scanln(&userInput)

	err := verifyPassword(userInput)
	for err != nil {
		fmt.Printf("\n\nTry again or control+c to quit: ")

		fmt.Scanln(&userInput)

		err = verifyPassword(userInput)
	}

}
