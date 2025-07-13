package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"♠ Spades", "♦ Diamonds", "♥ Hearts", "♣ Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// receiver function for deck type -- prints the deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// function to write the deck to a file

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(strings.Join([]string(d), ",")), 0666)
}

// function to read a deck from a file

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

// shuffle the deck of cards
func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
