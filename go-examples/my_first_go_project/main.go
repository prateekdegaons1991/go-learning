package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Rex Restaurant")
	items()

}

func items() string {

	fmt.Println("Enter the Items Ordered: ")
	fmt.Scan()
	return "Items Ordered"
}
