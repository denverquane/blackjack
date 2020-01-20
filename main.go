package main

import (
	"fmt"
	"github.com/denverquane/blackjack/deck"
	"math/rand"
)

const SEED = 1234567

func main() {
	rand.Seed(SEED)

	singleDeck := deck.MakeDeck()

	card1, _ := singleDeck.PullRandomCard()
	card2, _ := singleDeck.PullRandomCard()
	card3, _ := singleDeck.PullRandomCard()
	card4, _ := singleDeck.PullRandomCard()
	hand := deck.MakeHand(card1, card2)
	hand.Add(card3)
	hand.Add(card4)
	fmt.Println(hand.ToString())

}
