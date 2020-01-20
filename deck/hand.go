package deck

import "strconv"

type Hand struct {
	cards []Card
	//array for hard and soft values of the hand
	values []int
}

func MakeHand(card1 Card, card2 Card) *Hand {
	hand := Hand{
		cards:  make([]Card, 0),
		values: make([]int, 1),
	}
	hand.values[0] = 0
	hand.Add(card1)
	hand.Add(card2)
	return &hand
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

func (hand Hand) ToString() string {
	str := "Hand has cards:\n"
	for _, v := range hand.cards {
		str += "  " + v.ToString() + "\n"
	}
	str += "and values:\n"
	for _, v := range hand.values {
		str += "  " + strconv.FormatInt(int64(v), 10) + "\n"
	}
	return str
}
