package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(name string) error {
	return ioutil.WriteFile(name, []byte(d.toString()), 0666)
}

func newDeckFromFile(name string) deck {
	bs, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("Error >", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() deck {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range d {
		// Generate random
		// Swap card with random position
		n := r.Intn(len(d) - 1)

		d[i], d[n] = d[n], d[i]
	}
	return d
}
