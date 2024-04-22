package main

import "fmt"

func main() {

	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, i := range ints {
		if i%2 == 0 {
			fmt.Printf("%d Is Even number\n", i)
		} else {
			fmt.Printf("%d Is Odd number\n", i)
		}
	}
}
