package deck

import "log"

type Dealer struct {
	hitsSoft17 bool
	hand       Hand
}

func MakeDealerAndHand(hitsSoft17 bool, shoe *Shoe) Dealer {
	card1, err := shoe.PullRandomCard()
	if err != nil {
		log.Println(err)
	}
	return Dealer{hitsSoft17: hitsSoft17, hand: MakeHand(card1)}
}

func (dealer *Dealer) AddDownCard(card Card) {
	dealer.hand.Add(card)
}

func (dealer Dealer) UpCard() Card {
	return dealer.hand.FirstCard()
}

func (dealer Dealer) Hand() Hand {
	return dealer.hand
}

func (dealer Dealer) DoesHit() bool {
	if dealer.hand.values[0] < 17 || (dealer.hand.values[0] == 17 && dealer.hitsSoft17 && dealer.hand.IsSoft()) {
		return true
	} else {
		return false
	}
}

func (dealer *Dealer) Hit(shoe *Shoe) {
	card, err := shoe.PullRandomCard()
	if err != nil {
		log.Println(err)
	}
	dealer.hand.Add(card)
}
