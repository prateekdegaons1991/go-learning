package main

import "fmt"

// func getLanguages() (string, string) {
// 	return "Go", "Python"
// }

// func add(a, b int) int {
// 	return a + b
// }

func processIt(fn func(a int) int) {
	returnValue := fn(4)
	println(returnValue)
}

func main() {
	// result := add(1, 2)
	// fmt.Println(result)

	// language1, _ := getLanguages()
	// fmt.Println(language1)

	fn := func(a int) int {
		return a * 2
	}

	processIt(fn)

	// Variadic function call

	sum(1, 2, 3, 4, 5)
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}

// Variadic functions

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total

}



