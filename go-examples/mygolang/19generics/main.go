package main

import "fmt"

// This is a simple example of using generics in Go.

func printSlice[T int | string | any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// create struct having element of type T // generics
// This allows us to create a stack that can hold any type of elements.
type Stack[T any] struct {
	elements []T
}

func main() {
	fmt.Println("Let's learn, generics in Go!")

	// Using a slice of integers
	nums := []int{1, 2, 3, 4, 5}
	printSlice(nums)

	// Using a slice of strings
	names := []string{"apple", "banana", "cherry"}
	printSlice(names)

	// Using a slice of mixed types is not allowed in this example
	mixed := []any{"apple", 1, "banana", 2}
	printSlice(mixed) // This will not compile as mixed is not of type []T

	mystack := Stack[int]{
		elements: []int{1, 2, 3},
	}
	fmt.Println("Stack elements:", mystack.elements)

	mystack2 := Stack[string]{
		elements: []string{"apple", "banana", "cherry"},
	}
	fmt.Println("Stack elements:", mystack2.elements)
}
