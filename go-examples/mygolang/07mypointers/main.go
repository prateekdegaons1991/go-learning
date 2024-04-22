package main

import "fmt"

func main() {
	fmt.Println("Welcome to pinters")

	myNumber := 23

	var ptr = &myNumber

	*ptr = *ptr + 2 

	fmt.Println("Value of ptr is: ", *ptr)

}
