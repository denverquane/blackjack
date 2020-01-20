package deck

type Deck struct {
	Cards []Card
}

func MakeDeck() *Deck {
	deck := Deck{}
	deck.Cards = make([]Card, 52)
	count := 0

	for suit := SPADES; suit <= DIAMONDS; suit++ {
		for rank := ACE; rank <= KING; rank++ {
			deck.Cards[count] = Card{Rank: rank, Suit: suit}
			count++
		}
	}
	return &deck
}
