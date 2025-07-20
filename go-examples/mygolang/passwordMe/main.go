package main

import (
	"fmt"

	"github.com/prateekdegaons1991/passwordme/passwords"
)

var useUppercase, useDigits, useSpecialChars bool
var length int

func main() {
	fmt.Println("Welcome to PasswordMe!")
	UserInput()
	validChar := passwords.GeneratePassword(length, useUppercase, useDigits, useSpecialChars)
	fmt.Println("Valid characters for password generation:", validChar)

}

func UserInput() {
	fmt.Println("Please select the options for your password:")
	fmt.Print("Use uppercase letters? (Y/N): ")
	var upperCaseInput string
	fmt.Scanln(&upperCaseInput)
	switch upperCaseInput {
	case "Y", "y":
		useUppercase = true
	case "N", "n":
		useUppercase = false
	default:
		fmt.Println("Invalid input. Please enter Y or N.")
		return
	}

	fmt.Print("Use digits? (Y/N): ")
	var digitsInput string
	fmt.Scanln(&digitsInput)
	switch digitsInput {
	case "Y", "y":
		useDigits = true
	case "N", "n":
		useDigits = false
	default:
		fmt.Println("Invalid input. Please enter Y or N.")
		return
	}

	fmt.Print("Use numbers? (Y/N): ")
	var useSpecialCharsInput string
	fmt.Scanln(&useSpecialCharsInput)
	switch useSpecialCharsInput {
	case "Y", "y":
		useSpecialChars = true
	case "N", "n":
		useSpecialChars = false
	default:
		fmt.Println("Invalid input. Please enter Y or N.")
		return
	}

	fmt.Print("Enter desired password length: ")
	fmt.Scanln(&length)

	if length <= 0 {
		fmt.Println("Password length must be greater than 0.")
		return
	}
}
