package main

func main() {
	cards := parseDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
