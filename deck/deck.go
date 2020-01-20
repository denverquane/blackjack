package deck

import (
	"errors"
	"math/rand"
)

type Deck struct {
	cards []Card
}

func MakeDeck() *Deck {
	deck := Deck{}
	deck.cards = make([]Card, 52)
	count := 0

	for suit := SPADES; suit <= DIAMONDS; suit++ {
		for rank := ACE; rank <= KING; rank++ {
			deck.cards[count] = Card{Rank: rank, Suit: suit}
			count++
		}
	}
	return &deck
}

func (deck *Deck) CardsRemaining() int {
	return len(deck.cards)
}

func (deck *Deck) PullRandomCard() (Card, error) {
	if len(deck.cards) == 0 {
		return Card{NULLRANK, NULLSUIT}, errors.New("Deck is empty, cannot pull random card")
	}
	i := int(rand.Int31n(int32(len(deck.cards))))
	card := deck.cards[i]
	deck.cards = removeCard(deck.cards, i)
	return card, nil
}

func removeCard(s []Card, i int) []Card {
	//move the card to the end
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	//strip the last card off (much faster)
	return s[:len(s)-1]
}
