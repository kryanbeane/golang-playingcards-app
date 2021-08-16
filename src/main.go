package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Queen of Hearts")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
