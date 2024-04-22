package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")

	Pratik := User{"Pratik", "pratikdegaon@gmail.com", true, 24}
	fmt.Println(Pratik)
	fmt.Printf("Details are %+v\n", Pratik)
	fmt.Printf(Pratik.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
