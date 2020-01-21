package deck

import "strconv"

type Hand struct {
	cards []Card
	//array for hard and soft values of the hand
	values []int
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

func (hand Hand) IsSoft() bool {
	for _, v := range hand.cards {
		if v.Rank == ACE {
			return true
		}
	}
	return false
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
