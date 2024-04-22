package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	suites := []string{"Raja", "Rani", "Badam", "Kavva"}
	values := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range suites {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
