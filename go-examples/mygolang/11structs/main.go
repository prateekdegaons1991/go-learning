package main

import "fmt"

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
	Name   string
	Email  string
	Status bool
	Age    int
}
