package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, instead got %v", len(d))
	}

	if d[0] != "2 of Clubs" {
		t.Errorf("Expected the first entry to be '2 of Clubs', instead got %v", d[0])
	}

	if d[len(d)-1] != "Ace of Spades" {
		t.Errorf("Expected the first entry to be 'Ace of Spades', instead got %v", d[0])
	}
}

func TestSaveToFileAndParseDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	parsedDeck := parseDeckFromFile("_decktesting")

	if len(parsedDeck) != 52 {
		t.Errorf("Expected parsedDeck to have length 52, but instead got %v", len(parsedDeck))
	}

	os.Remove("_decktesting")
}
