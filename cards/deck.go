package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of deck
// which is a slice of string

type deck []string

// no receiver
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "hearts", "Club"}
	cardValues := []string{"Ace", "two", "three", "four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// append returns a new cards
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

// any data of type deck gets access to print function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// type conversion of deck to slice of string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// log error and quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// convert bs to a string
	// split string by ","
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
