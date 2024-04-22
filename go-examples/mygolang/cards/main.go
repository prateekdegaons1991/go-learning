package main

import "fmt"

func main() {
	cards := newDeckFromFile("MyCards")

	// hand, reminingCards := deal(cards, 5)

	// hand.print()
	// reminingCards.print()
	fmt.Println(cards)
	// cards.saveToFile("MyCards2")
	fmt.Println(deal(cards, 5))
}
