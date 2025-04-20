package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Structs in golang")

	Pratik := User{"Pratik", "pratikdegaon@gmail.com", true, 24}
	var Eshita *User = &User{"Eshita", "eshita@gmail.com", true, 23}

	fmt.Println(*Eshita)
	fmt.Println(Pratik)

	Pratik.Name = "Pratik English"

	fmt.Println(Pratik)

	// lets create pointer to struct

	var PratikPointer *User
	PratikPointer = &Pratik
	PratikPointer.Name = "Pratik Degaonkar"
	fmt.Println(*PratikPointer)

}

// Capital letters are exported
// Structs are used to create a collection of fields
type User struct {
	Name      string
	LastName  string
	Email     string
	Status    bool
	Age       string
	createdAt time.Time
}

type UserAddress struct {
	HouseNumber int
	Street      string
	City        string
	State       string
	Pincode     int
}

// nested structs
// You can also create nested structs, which are structs that contain other structs as fields.
// This allows you to create complex data structures that can represent real-world entities.
// For example, you can create a struct that represents a user and another struct that represents the user's address.
// You can then create a third struct that contains both the user and the address structs.
type UserDetails struct {
	User
	UserAddress
}

// Methods
// In Go, you can define methods on structs, which are functions that operate on the struct's fields.
// Methods are defined using the func keyword, followed by the receiver type (the struct type) and the method name.
// The receiver type is specified in parentheses before the method name.
// Methods can be used to encapsulate behavior and provide a way to interact with the struct's data.
// For example, you can define a method on the User struct that returns the user's full name.

func (u User) isUserActive() bool {
	return u.Status
}

func getUserData(promtText string) (string, error) {
	fmt.Print(promtText)
	var value string
	fmt.Scanln(&value)
	if value == "" {
		return "", fmt.Errorf("value cannot be empty")
	}
	return value, nil
}
