package main

import (
	"fmt"
	"time"
)

// Structs are used to group related data together.
// They are similar to classes in other programming languages but are simpler.
// Define a struct
type user struct {
	firsrtName string
	lastName   string
	birthDate  string
	createdAt  time.Time
}

func main() {
	fmt.Println("Structs in Go")

	var appUser user
	appUser.firsrtName = "John"
	appUser.lastName = "Doe"
	appUser.birthDate = "01/01/1990"
	appUser.createdAt = time.Now().Truncate(time.Hour * 24)
	fmt.Println(appUser)
}
