package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
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

//read data from hard drive
func newDeckFromFile(filename string) deck {
	bs, err:= ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	ss := strings.Split(string(bs), ",") // ace of spades, tow of spades, three of spades,
	return deck(ss)
}

//suffle data
func (d deck) shuffle() {
	//generator random number
	source:= rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}