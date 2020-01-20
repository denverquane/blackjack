package deck

type CardRank int
type CardSuit int

const (
	NULLRANK CardRank = iota
	ACE
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

var CardRankStrings = map[CardRank]string{
	NULLRANK: "Null",
	ACE:      "Ace",
	TWO:      "Two",
	THREE:    "Three",
	FOUR:     "Four",
	FIVE:     "Five",
	SIX:      "Six",
	SEVEN:    "Seven",
	EIGHT:    "Eight",
	NINE:     "Nine",
	TEN:      "Ten",
	JACK:     "Jack",
	QUEEN:    "Queen",
	KING:     "King",
}

const (
	NULLSUIT CardSuit = iota
	SPADES
	HEARTS
	CLUBS
	DIAMONDS
)

var CardSuitStrings = map[CardSuit]string{
	NULLSUIT: "Null",
	SPADES:   "Spades",
	HEARTS:   "Hearts",
	CLUBS:    "Clubs",
	DIAMONDS: "Diamonds",
}

type Card struct {
	Rank CardRank
	Suit CardSuit
}

func (card Card) ToString() string {
	return CardRankStrings[card.Rank] + " of " + CardSuitStrings[card.Suit]
}
