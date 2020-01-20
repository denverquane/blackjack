package main

import (
	"fmt"
	"github.com/denverquane/blackjack/deck"
)

const SEED = "ghf"

func main() {
	singleDeck := deck.MakeDeck()
	for _, v := range singleDeck.Cards {
		fmt.Println(v.ToString())
	}
}
