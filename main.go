package main

func main() {
	cards := newDeck()

	//every variable declared in a slice loop MUST be used in the callback function
	cards.print()
}
