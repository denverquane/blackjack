package deck

type Rules struct {
	dealerHitsSoft17    bool
	canDoubleAfterSplit bool
	shoeSize            int
}

func MakeRules(dealerHits, canDouble bool, shoeSize int) Rules {
	return Rules{
		dealerHitsSoft17:    dealerHits,
		canDoubleAfterSplit: canDouble,
		shoeSize:            shoeSize,
	}
}

func (rules Rules) DoesDealerHitSoft17() bool {
	return rules.dealerHitsSoft17
}
