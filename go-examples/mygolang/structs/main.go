package main

import "fmt"

var users = []string{}

type contactInfo struct {
	email string
	phone string
}

type Person struct {
	firstName string
	lastName  string
	age       int
	email     string
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
		},
	}

	p.updateName("Eshita")
	p.Print()
}
func (pointerToPerson *Person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p Person) Print() {
	fmt.Printf("%+v\n", p)
}
