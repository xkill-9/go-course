package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	f := d[0]
	if f != "Ace of Spades" {
		t.Errorf("Expected first to be Ace of Spades but got %v", f)
	}

	l := d[len(d)-1]
	if l != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %v", l)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktest")

	d := newDeck()
	d.saveToFile("_decktest")

	ld := newDeckFromFile("_decktest")

	if len(ld) != 16 {
		t.Errorf("Expected 16 cards, but got %v", len(ld))
	}

	os.Remove("_decktest")
}
