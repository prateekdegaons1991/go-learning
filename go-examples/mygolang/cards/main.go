package main

import "fmt"

func main() {
	// Create a new deck of cards
	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("all_cards")
	loadCards := newDeckFromFile("all_cards")
	fmt.Println(len(loadCards), "cards in the deck")

}
