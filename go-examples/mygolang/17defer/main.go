package main

import "fmt"

func main() {

	fmt.Println("Defer in Go")
	defer fmt.Println("One")
	fmt.Println("Two")

	MyDefer()
}

// defer loop

func MyDefer() {
	for i := range 5 {
		defer fmt.Print(i)
	}
 
}
