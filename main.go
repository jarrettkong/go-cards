package main

func main() {
	cards := deck{"Ace of Spades", "Five of Diamonds"}
	cards = append(cards, "Joker")

	cards.print()
}
