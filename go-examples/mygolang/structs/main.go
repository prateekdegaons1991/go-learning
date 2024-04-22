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

	alex := Person{
		firstName: "Pratik",
		lastName:  "Degaon",
		age:       32,
		email:     "pratikdegaon@gmail.com",
		contactInfo: contactInfo{
			email: "pratikdegaon@gmail.com",
			phone: "7411455742",
		},
	}

	alex.updateName("Eshita")
	alex.Print()

}
func (p Person) updateName(newName string) {
	p.firstName = newName
}
func (p Person) Print() {
	fmt.Printf("%v\n", p)
}
