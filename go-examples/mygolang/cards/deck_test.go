package main

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// test to see the deck has 52 cards

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}
}

// test to see if the first card is Ace of Spades

func TestFirstCard(t *testing.T) {
	deck := newDeck()
	if deck[0] != "Ace of ♠ Spades" {
		t.Errorf("Expected first card to be 'Ace of ♠ Spades', but got %v", deck[0])
	}
}

// test to see if the last card is King of Clubs
func TestLastCard(t *testing.T) {
	deck := newDeck()
	if deck[len(deck)-1] != "King of ♣ Clubs" {
		t.Errorf("Expected last card to be 'King of ♣ Clubs', but got %v", deck[len(deck)-1])
	}
}

// test to see if the deck can be saved to a file and read back correctly
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected loaded deck length of 52, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}

// write testcase using ginkgo and gomega to test if deck has 52 cards

var _ = Describe("Deck", func() {
	Context("NewDeck", func() {
		It("should have 52 cards", func() {
			deck := newDeck()
			Expect(len(deck)).To(Equal(52))
		})
	})
})
