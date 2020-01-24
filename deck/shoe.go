package deck

import (
	"errors"
	"log"
	"math"
	"math/rand"
)

type Shoe struct {
	decks      []Deck
	totalDecks int
	//number of card remaining in the deck before a shuffle
	penetrationShuffle int
	cardsRemaining     int

	count int
}

func MakeShoe(numDecks int) Shoe {
	shoe := Shoe{
		decks:              make([]Deck, numDecks),
		totalDecks:         numDecks,
		penetrationShuffle: rand.Intn(42) + 10,
		cardsRemaining:     numDecks * 52,
	}
	for i := 0; i < numDecks; i++ {
		shoe.decks[i] = MakeDeck()
	}
	return shoe
}

func (shoe Shoe) TrueCount() int {
	return shoe.count / int(math.Ceil(float64(shoe.cardsRemaining)/52.0))
}

func (shoe *Shoe) PullAndAddToHand(hand *Hand) {
	card, err := shoe.PullRandomCard()
	if err != nil {
		log.Println(err)
	}
	hand.Add(card)
}

func (shoe *Shoe) PullRandomCard() (Card, error) {
	if len(shoe.decks) > 0 && shoe.cardsRemaining > shoe.penetrationShuffle {
		i := rand.Intn(len(shoe.decks))
		if shoe.decks[i].CardsRemaining() > 0 {
			card, _ := shoe.decks[i].PullRandomCard()

			if card.Rank == TWO || card.Rank == THREE || card.Rank == FOUR || card.Rank == FIVE || card.Rank == SIX {
				shoe.count++
			} else if card.Rank == TEN || card.Rank == JACK || card.Rank == QUEEN || card.Rank == KING || card.Rank == ACE {
				shoe.count--
			}

			shoe.cardsRemaining--
			return card, nil

		} else {
			shoe.decks = removeDeck(shoe.decks, i)
			return shoe.PullRandomCard()
		}
	}

	return Card{}, errors.New("Shoe needs to be reshuffled")

}

func (shoe Shoe) GetCardsRemainingBeforeCut() int {
	return shoe.cardsRemaining - shoe.penetrationShuffle
}

func removeDeck(s []Deck, i int) []Deck {
	//move the card to the end
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	//strip the last card off (much faster)
	return s[:len(s)-1]
}
