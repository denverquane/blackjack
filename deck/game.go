package deck

type Rules struct {
	dealerHitsSoft17    bool
	canDoubleAfterSplit bool
	shoeSize            int
	blackjackPayout		float64
	maxBetSpread int
}

func MakeRules(dealerHits, canDouble bool, shoeSize int, blackjackPayout float64, betSpread int) Rules {
	return Rules{
		dealerHitsSoft17:    dealerHits,
		canDoubleAfterSplit: canDouble,
		shoeSize:            shoeSize,
		blackjackPayout:blackjackPayout,
		maxBetSpread:betSpread,
	}
}

func (rules Rules) BlackjackPayout() float64 {
	return rules.blackjackPayout
}

func (rules Rules) DoesDealerHitSoft17() bool {
	return rules.dealerHitsSoft17
}

func (rules Rules) MaxBetSpread() int {
	return rules.maxBetSpread
}
