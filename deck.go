package main

import (
	"fmt"
	"math/rand"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handsize int) (deck, deck) {
	return d[handsize:], d[:handsize]
}

func shuffle(d deck) deck {
	shuffle := deck{}

	for len := len(d); len > 0; len-- {
		card := rand.Intn(len)
		shuffle = append(shuffle, d[card])
		d[card] = d[len-1]
	}

	return shuffle
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
