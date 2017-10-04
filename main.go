package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	cards = shuffle(cards)
	cards, hand := deal(cards, 3)
	fmt.Println("Hand")
	hand.print()

	fmt.Println("Deck")
	cards.print()
}
