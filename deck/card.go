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

var CardRankRunes = map[CardRank]string{
	NULLRANK: "N",
	ACE:      "A",
	TWO:      "2",
	THREE:    "3",
	FOUR:     "4",
	FIVE:     "5",
	SIX:      "6",
	SEVEN:    "7",
	EIGHT:    "8",
	NINE:     "9",
	TEN:      "10",
	JACK:     "J",
	QUEEN:    "Q",
	KING:     "K",
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

var CardSuitRunes = map[CardSuit]string{
	NULLSUIT: "0",
	SPADES:   "♠",
	HEARTS:   "♥",
	CLUBS:    "♣",
	DIAMONDS: "♦",
}

type Card struct {
	Rank CardRank
	Suit CardSuit
}

func (card Card) ToString() string {
	return CardRankStrings[card.Rank] + " of " + CardSuitStrings[card.Suit]
}

func (card Card) ToAscii() string {
	return CardRankRunes[card.Rank] + " " + CardSuitRunes[card.Suit]
}
