package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of deck which will be slide of strings

type deck []string

func (d deck) print() {
	for index, value := range d {

		fmt.Println(index, value)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuit := []string{"Hearts", "Spades", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suit := range cardSuit {
		for _, value := range cardValues {
			cards = append(cards, value+" "+"of"+" "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) savetoFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckfromFile(fileName string) deck {

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	s := strings.Split(string(data), ",")
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
