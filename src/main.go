package main

func main() {
	cards := newDeckFromFile("my_cards")
	// cards2 := newDeckFromFile("my_cards2") // throws error: file not found
	cards.print()
}
