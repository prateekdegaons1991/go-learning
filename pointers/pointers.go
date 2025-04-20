package main

import "fmt"

func main() {
	age := 33 // Regular variable

	agePointer := &age // Pointer variable created with &

	fmt.Println("age:", *agePointer)

	fmt.Println("Years as an adult:", getAdultYears(&age))

}

func getAdultYears(age *int) int {
	return *age - 18
}
