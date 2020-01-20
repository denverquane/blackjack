package deck

type Dealer struct {
	hitsSoft17 bool
}

func MakeDealer(hitsSoft17 bool) Dealer {
	return Dealer{hitsSoft17: hitsSoft17}
}

func (dealer Dealer) DoesHit(hand Hand) bool {
	if hand.values[0] < 17 || (hand.values[0] == 17 && dealer.hitsSoft17 && hand.IsSoft()) {
		return true
	} else {
		return false
	}
}
