package deck

type StrategyAction byte

const (
	STND                   = StrategyAction(STAND)
	HIIT                   = StrategyAction(HIT)
	SPLT                   = StrategyAction(SPLIT)
	SPLTDBL StrategyAction = 'G' //Split if doubling after split is allowed
	HITDBL  StrategyAction = 'I' //Double if allowed, otherwise hit
	STNDDBL StrategyAction = 'J' //Double if allowed, otherwise stand
)

func PlayerStrategy(gameRules Rules, playerHand Hand, dealerCard Card) PlayerAction {
	isSoft := playerHand.IsSoft()
	isSplit := playerHand.IsSplit()
	value := playerHand.HighestPlay()
	strategy := STND
	playerIdx := 0
	dealerIdx := int(dealerCard.Rank)
	if dealerCard.Rank == JACK || dealerCard.Rank == QUEEN || dealerCard.Rank == KING {
		dealerIdx = int(TEN)
	} else if dealerCard.Rank == ACE {
		dealerIdx = 11
	}
	dealerIdx = dealerIdx - 2
	//log.Printf("Dealer index: %d\n", dealerIdx)

	if playerHand.CanSplit() {
		//log.Println("Hand is Pairs")
		playerIdx = value / 2
		playerIdx = playerIdx - 2
		strategy = DealerHitsSoft17PairsStrategy[playerIdx][dealerIdx]
	} else if isSoft {
		//log.Println("Hand is soft")
		if value > 19 {
			value = 19
		} else if value < 13 {
			value = 13
		}
		playerIdx = value - 13 //offset for soft hands
		strategy = DealerHitsSoft17SoftStrategy[playerIdx][dealerIdx]
	} else {
		if value > 17 {
			value = 17
		}
		playerIdx = value - 8
		if playerIdx < 0 {
			playerIdx = 0
		}
		strategy = DealerHitsSoft17Strategy[playerIdx][dealerIdx]
	}
	//log.Printf("Player index: %d\n", playerIdx)

	switch strategy {
	case STND:
		return STAND
	case HIIT:
		return HIT
	case SPLT:
		return SPLIT
	case SPLTDBL:
		if gameRules.canDoubleAfterSplit {
			return SPLIT
		} else {
			return HIT
		}
	case HITDBL:
		//if the hand is already split, and we're allowed to double after a split
		if isSplit && gameRules.canDoubleAfterSplit {
			return DOUBLE
		} else {
			return HIT
		}
	case STNDDBL:
		//if the hand is already split, and we're allowed to double after a split
		if isSplit && gameRules.canDoubleAfterSplit {
			return DOUBLE
		} else {
			return STAND
		}
	}
	return NULLACTION
}

var DealerHitsSoft17Strategy = [10][10]StrategyAction{
	// Dealer Card ->
	// 2       3       4       5       6       7       8       9       10      A      Player
	{HIIT, HIIT, HIIT, HIIT, HIIT, HIIT, HIIT, HIIT, HIIT, HIIT},                   //4-8
	{HIIT, HITDBL, HITDBL, HITDBL, HITDBL, HIIT, HIIT, HIIT, HIIT, HIIT},           //9
	{HITDBL, HITDBL, HITDBL, HITDBL, HITDBL, HITDBL, HITDBL, HITDBL, HIIT, HIIT},   //10
	{HITDBL, HITDBL, HITDBL, HITDBL, HITDBL, HITDBL, HITDBL, HITDBL, HITDBL, HIIT}, //11
	{HIIT, HIIT, STND, STND, STND, HIIT, HIIT, HIIT, HIIT, HIIT},                   //12
	{STND, STND, STND, STND, STND, HIIT, HIIT, HIIT, HIIT, HIIT},                   //13
	{STND, STND, STND, STND, STND, HIIT, HIIT, HIIT, HIIT, HIIT},                   //14
	{STND, STND, STND, STND, STND, HIIT, HIIT, HIIT, HIIT, HIIT},                   //15
	{STND, STND, STND, STND, STND, HIIT, HIIT, HIIT, HIIT, HIIT},                   //16
	{STND, STND, STND, STND, STND, STND, STND, STND, STND, STND},                   //17+
}

var DealerHitsSoft17SoftStrategy = [7][10]StrategyAction{
	// Dealer Card ->
	// 2       3       4       5       6       7       8       9       10      A      Player
	{HIIT, HIIT, HIIT, HITDBL, HITDBL, HIIT, HIIT, HIIT, HIIT, HIIT},         //S13
	{HIIT, HIIT, HIIT, HITDBL, HITDBL, HIIT, HIIT, HIIT, HIIT, HIIT},         //S14
	{HIIT, HIIT, HITDBL, HITDBL, HITDBL, HIIT, HIIT, HIIT, HIIT, HIIT},       //S15
	{HIIT, HIIT, HITDBL, HITDBL, HITDBL, HIIT, HIIT, HIIT, HIIT, HIIT},       //S16
	{HIIT, HITDBL, HITDBL, HITDBL, HITDBL, HIIT, HIIT, HIIT, HIIT, HIIT},     //S17
	{STND, STNDDBL, STNDDBL, STNDDBL, STNDDBL, STND, STND, HIIT, HIIT, HIIT}, //S18
	{STND, STND, STND, STND, STND, STND, STND, STND, STND, STND},             //S19+
}

var DealerHitsSoft17PairsStrategy = [10][10]StrategyAction{
	// Dealer Card ->
	// 2       3       4       5       6       7       8      9       10       A      Player
	{SPLTDBL, SPLTDBL, SPLT, SPLT, SPLT, SPLT, HIIT, HIIT, HIIT, HIIT},           //2,2
	{SPLTDBL, SPLTDBL, SPLT, SPLT, SPLT, SPLT, HIIT, HIIT, HIIT, HIIT},           //3,3
	{HIIT, HIIT, HIIT, SPLTDBL, SPLTDBL, HIIT, HIIT, HIIT, HIIT, HIIT},           //4,4
	{HITDBL, HITDBL, HITDBL, HITDBL, HITDBL, HITDBL, HITDBL, HITDBL, HIIT, HIIT}, //5,5
	{SPLTDBL, SPLT, SPLT, SPLT, SPLT, HIIT, HIIT, HIIT, HIIT, HIIT},              //6,6
	{SPLT, SPLT, SPLT, SPLT, SPLT, SPLT, HIIT, HIIT, HIIT, HIIT},                 //7,7
	{SPLT, SPLT, SPLT, SPLT, SPLT, SPLT, SPLT, SPLT, SPLT, SPLT},                 //8,8
	{SPLT, SPLT, SPLT, SPLT, SPLT, STND, SPLT, SPLT, STND, STND},                 //9,9
	{STND, STND, STND, STND, STND, STND, STND, STND, STND, STND},                 //10,10
	{SPLT, SPLT, SPLT, SPLT, SPLT, SPLT, SPLT, SPLT, SPLT, SPLT},                 //A,A
}
