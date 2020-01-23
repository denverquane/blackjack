package deck

import "strconv"

type Hand struct {
	cards []Card
	//array for hard and soft values of the hand
	values []int
	split  bool
}

func MakeHand(card1 Card) Hand {
	hand := Hand{
		cards:  make([]Card, 0),
		values: make([]int, 1),
	}
	hand.values[0] = 0
	hand.Add(card1)
	return hand
}

func (hand *Hand) Add(card Card) {
	if card.Rank == ACE {
		orig := hand.values[0]
		for i, v := range hand.values {
			hand.values[i] = v + 1
		}
		hand.values = append(hand.values, orig+11)
	} else if card.Rank == JACK || card.Rank == QUEEN || card.Rank == KING {
		for i, v := range hand.values {
			hand.values[i] = v + 10
		}
	} else {
		for i, v := range hand.values {
			hand.values[i] = v + int(card.Rank)
		}
	}
	hand.cards = append(hand.cards, card)
}

func (hand Hand) Split() (h1, h2 Hand) {
	h1 = MakeHand(hand.cards[0])
	h1.split = true
	h2 = MakeHand(hand.cards[1])
	h2.split = true
	return h1, h2
}

func (hand Hand) CanSplit() bool {
	return len(hand.cards) == 2 && hand.cards[0].Rank == hand.cards[1].Rank
}

func (hand Hand) IsSoft() bool {
	for _, v := range hand.cards {
		if v.Rank == ACE {
			return true
		}
	}
	return false
}

func (hand Hand) IsSplit() bool {
	return hand.split
}

func (hand Hand) IsBust() bool {
	return hand.values[0] > 21
}

func (hand Hand) FirstCard() Card {
	return hand.cards[0]
}

func (hand Hand) HighestPlay() int {
	for i := len(hand.values) - 1; i > -1; i-- {
		if hand.values[i] <= 22 {
			return hand.values[i]
		}
	}
	//all busts
	return 0
}

func (hand Hand) HasBlackjack() bool {
	card1 := hand.cards[0].Rank
	card2 := hand.cards[1].Rank
	return ((card1 == JACK || card1 == QUEEN || card1 == KING || card1 == TEN) && card2 == ACE) ||
		((card2 == JACK || card2 == QUEEN || card2 == KING || card2 == TEN) && card1 == ACE)
}

func (hand Hand) ToString(printValues bool) string {
	str := "Hand:\n"
	for _, v := range hand.cards {
		str += "  " + v.ToString() + "\n"
	}
	if printValues {
		str += "and values:\n"
		for _, v := range hand.values {
			str += "  " + strconv.FormatInt(int64(v), 10) + "\n"
		}
	}
	return str
}
func (hand Hand) ToAscii(printValues bool) string {
	str := "Hand:\n"
	for _, v := range hand.cards {
		str += "  " + v.ToAscii()
	}
	if printValues {
		str += "\nand values:\n"
		for _, v := range hand.values {
			str += "  " + strconv.FormatInt(int64(v), 10) + "\n"
		}
	}
	return str
}
