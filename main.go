package main

import "fmt"

func main() {
	cards := newDeck()
	hand, cards := cards.deal(4)
	// cards.print()
	// hand.print()
	fmt.Println(hand.toString())
	fmt.Println(cards.toString())
}
