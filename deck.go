package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
//which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
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
	//these return values point to sections of the slice, they do not alter the original deck.
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ", ")
	return deck(s)
}

func (d deck) shuffle() {
	// rand is pseudo random and needs a different seed to truly generate a random number.
	// use time to generate a unique time and pass it as seed
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
