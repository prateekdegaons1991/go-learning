package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in Go!")
	// Declare and initialize a slice
	var fruits = []string{"apple", "banana", "peach"}
	fmt.Println("Fruits: ", fruits)

	fruits = append(fruits, "mango", "orange")
	fmt.Println("Fruits after append: ", fruits)

	fruits = append(fruits[:3], fruits[3:]...) // Remove the first element
	fmt.Println("Fruits after removing first element: ", fruits)

	fruits = append(fruits[:2], fruits[3:]...) // Remove the third element
	fmt.Println("Fruits after removing third element: ", fruits)

	// Declare and initialize a slice with make
	highScores := make([]int, 5)
	highScores[0] = 100
	highScores[1] = 99
	highScores[2] = 300
	highScores[3] = 400
	highScores[4] = 500

	fmt.Println("High Scores: ", highScores)
	highScores = append(highScores, 10, 00)
	fmt.Println("High Scores after append: ", highScores)
	// Sort the slice
	sort.Ints(highScores)
	fmt.Println("High Scores after sorting: ", highScores)
	// Declare and initialize a slice with make and a length of 0
	emptySlice := make([]int, 0)
	fmt.Println("Empty Slice: ", emptySlice)
	// Declare and initialize a slice with make and a length of 0 and a capacity of 5
	emptySliceWithCap := make([]int, 0, 5)
	fmt.Println("Empty Slice with capacity: ", emptySliceWithCap)
	// Declare and initialize a slice with make and a length of 0 and a capacity of 5
	emptySliceWithCapAndLen := make([]int, 5, 5)
	fmt.Println("Empty Slice with capacity and length: ", emptySliceWithCapAndLen)

	//how to remove a value from a slices based on index

	var courses = []string{"reactjs", "angularjs", "vuejs", "javascript", "python"}
	fmt.Println("Courses: ", courses)
	var indexToRemove int = 2
	courses = append(courses[:indexToRemove], courses[indexToRemove+1:]...)
	fmt.Println("Courses after removing index 2: ", courses)

	//how to remove a value from a slices based on value
	var courses2 = []string{"reactjs", "angularjs", "vuejs", "javascript", "python"}
	fmt.Println("Courses2: ", courses2)
	var valueToRemove string = "python"
	var indexToRemove2 int
	for i, v := range courses2 {
		if v == valueToRemove {
			indexToRemove2 = i
			break
		}
	}
	courses2 = append(courses2[:indexToRemove2], courses2[indexToRemove2+1:]...)
	fmt.Printf("Courses2 after removing %v: %v\n", valueToRemove, courses2)

}
