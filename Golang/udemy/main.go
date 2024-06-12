package main

func main() {
	//save file
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	//read file
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards:= newDeck()
	cards.shuffle()
	cards.print()
}
