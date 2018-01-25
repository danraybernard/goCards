package main

import (
	"fmt"
)

// create a new type of 'deck'
//which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Threee", "Four"}
	//every variable declared in a slice loop MUST be used in the callback function. Use _ to negate the index
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//d is referencing the individual deck that the function is acting on
// golang convention dictates that the receiver variable should be one or two letters related to the original type. So d from deck.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//the second parens tell what the function is returning. In this case two decks, one drawn from the original, and one with those cards subtracted
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
