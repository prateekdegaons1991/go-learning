package main

import "fmt"

func main() {
	fmt.Println("Fucntions in Go")
	greater()

	result := proAdder(1, 2, 3, 4, 5)
	fmt.Println("Sum of numbers is:", result)

}

func greater() {
	fmt.Println("Namaste from greater function")
}

// function with parameters
func add(a int, b int) int {
	return a + b
}

// function with multiple parameters
func add2(a, b int) int {
	return a + b
}

// function with multiple parameters and return values
func addSub(a, b int) (int, int) {
	return a + b, a - b
}

// proAdder is a function that takes a variable number of integer arguments and returns their sum.
func proAdder(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
