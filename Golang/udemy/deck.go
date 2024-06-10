package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// function reciver
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _,suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i,card := range d {
		fmt.Println(i,card)
	}
}

//slice function
func deal(d deck, handleSize int) (deck, deck) {
	return d[:handleSize], d[handleSize:]
}

//string joins
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//save to file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0755)
}