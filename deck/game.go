package deck

type Rules struct {
	dealerHitsSoft17    bool
	canDoubleAfterSplit bool
	shoeSize            int
	blackjackPayout     float64
	maxBetSpread        int
	resplit             bool
}

func MakeRules(dealerHits, canDouble bool, shoeSize int, blackjackPayout float64, betSpread int, resplit bool) Rules {
	return Rules{
		dealerHitsSoft17:    dealerHits,
		canDoubleAfterSplit: canDouble,
		shoeSize:            shoeSize,
		blackjackPayout:     blackjackPayout,
		maxBetSpread:        betSpread,
		resplit:             resplit,
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

func (rules Rules) Resplit() bool {
	return rules.resplit
}
