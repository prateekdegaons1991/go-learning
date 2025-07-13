package main

func main() {
	// Create a new deck of cards
	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("all_cards")
	loadCards := newDeckFromFile("all_cards")
	loadCards.print()

	// Deal 5 cards from the deck
	hand, remainingDeck := deal(cards, 5)
	hand.saveToFile("hand_cards")
	remainingDeck.saveToFile("remaining_deck")

}
