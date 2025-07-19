package main

import (
	"fmt"

	"github.com/prateekdegaons1991/passwordme/passwords"
)

func main() {
	fmt.Println("Welcome to PasswordMe!")
	validate := passwords.ValidChars(true, true, true)
	fmt.Println("Generated valid characters for password:", validate)
	password := passwords.GeneratePassword(12, true, true, true)
	fmt.Println("Generated password:", password)
	fmt.Println("Thank you for using PasswordMe!")
}

// This is a simple Go program that prints a welcome message to the console.
