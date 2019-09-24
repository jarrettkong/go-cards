package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) deal(size int) (deck, deck) {
	return d[:size], d[size:]
}
