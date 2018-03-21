package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type card struct {
	name  string
	suit  string
	rank  string
	value int
}

type deck []card

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardRanks := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for offset, rank := range cardRanks {
			cards = append(cards, card{
				name:  rank + " of " + suit,
				suit:  suit,
				rank:  rank,
				value: (13 - offset)})
		}
	}
	return cards
}

func deal(d deck, handsize int) (deck, deck) {
	return d[handsize:], d[:handsize]
}

func shuffle(d deck) deck {
	realRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	shuffle := deck{}

	for len := len(d); len > 0; len-- {
		card := realRand.Intn(len)
		shuffle = append(shuffle, d[card])
		d[card] = d[len-1]
	}

	return shuffle
}

func (d deck) print() {
	fmt.Println(d.toString())
}

func (d deck) isFlush() bool {
	if len(d) == 0 {
		return false
	}

	firstSuit := d[0].suit

	for _, card := range d {
		if card.suit != firstSuit {
			return false
		}
	}

	return true
}

func (d deck) toString() string {
	names := []string{}
	for _, card := range d {
		names = append(names, card.name)
	}
	return strings.Join(names, "\n")
}
