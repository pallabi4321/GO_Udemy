package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck of length 16, but got %v", len(d))
	}

	if d[0] != "Ace of Diamonds" {
		t.Errorf("Expected Ace of Diamonds, but got %v", d[0])
	}

	if d[ len(d) - 1 ] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	deckloaded := newDeckFromFile("_decktesting")

	if len(deckloaded) != 16 {
		t.Errorf("Expected 16, got %v", len(deckloaded))
	}

	os.Remove("_decktesting")
}