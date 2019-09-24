package main

import "fmt"

func main() {
	cards := newDeck()
	hand, cards := cards.deal(4)
	cards.print()
	fmt.Println(hand)
}
