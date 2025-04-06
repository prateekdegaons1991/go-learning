package main

import (
	"fmt"
	"strconv"
	"structs/user"
)

var users = []string{}

func main() {

	firstName := user.GetUserData("Enter your first name: ")
	lastName := user.GetUserData("Enter your lastname: ")
	email := user.GetUserData("Enter your email: ")
	ageInput := user.GetUserData("Enter your age: ")
	age, err := strconv.Atoi(ageInput)
	if err != nil {
		fmt.Println("Invalid age input. Please enter a valid number.")
		return
	}
	phone := user.GetUserData("Enter your phone: ")

	// Create a new person using the constructor

	newUser, err := user.NewPerson(firstName, lastName, age, email, phone)

	newUser.OutputUserDetails()

	// p := Person{
	// 	firstName: "Pratik",
	// 	lastName:  "Degaon",
	// 	age:       32,
	// 	createdAt: time.Now(),
	// 	contactInfo: contactInfo{
	// 		email: "pratikdegaon@gmail.com",
	// 		phone: "7411455742",
	// 	},
	// }

	// p.updateName("Eshita")
	// p.Print()
}

// This method will update the name of the person
// It uses a pointer receiver to modify the original struct
// func (pointerToPerson *Person) updateName(newName string) { // Pointer receiver
// 	// This method will update the name of the person
// 	(*pointerToPerson).firstName = newName // Dereferencing the pointer to access the field
// }

// func (p Person) Print() {
// 	fmt.Printf("%+v\n", p)
// }
