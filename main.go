package main

func main() {
	cards := newDeck()
	hand, cards := cards.deal(4)
	cards.print()
	hand.print()
	cards.saveToFile("my_cards")
}
