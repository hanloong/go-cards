package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck to have 52 cards but got: %v", len(d))
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	d, hand := deal(d, 3)

	if len(hand) != 3 {
		t.Errorf("Expected hand to have 3 cards but got: %v", len(hand))
	}
}
