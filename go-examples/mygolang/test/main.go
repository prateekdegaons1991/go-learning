package main

import "fmt"

func main() {
	fmt.Println("Lets Learn about receiever and custome type")

	cards := newDeck()

	cards.print()
}
