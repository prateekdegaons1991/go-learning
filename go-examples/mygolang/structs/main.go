package main

import "fmt"

var users = []string{}

type contactInfo struct {
	email string
	phone string
	zip   int
}

// embedding structs
// Person struct with embedded contactInfo struct
type Person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
}

func main() {

	p := Person{
		firstName: "Pratik",
		lastName:  "Degaon",
		age:       32,
		contactInfo: contactInfo{
			email: "pratikdegaon@gmail.com",
			phone: "7411455742",
			zip:   411057,
		},
	}

	p.updateName("Eshita", "Mahant")
	p.Print()

	var alex Person // another way to create a struct
	alex.firstName = "Alexy"
	alex.lastName = "Anderson"
	alex.age = 30
	// pointer to the struct // but not necessary to use pointer receiver , instead of pointer receiver we can use value receiver like below
	alex.updateName("Alexander", "Anderson")

	fmt.Printf("%v's age is %v\n", alex.firstName, alex.age)
}

// updateName method to change the first name of a Person , this is a receiver function
func (pointerToPerson *Person) updateName(newName, newLastName string) {
	(*pointerToPerson).firstName = newName // updating first name
	pointerToPerson.lastName = newLastName // updating last name as well
}

func (p Person) Print() {
	fmt.Printf("%+v\n", p)
}
