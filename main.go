package main

import (
	"fmt"
	"github.com/denverquane/blackjack/deck"
	"log"
	"math/rand"
	"time"
)

const SEED = 1234567

//Dumb heuristic, assume the last 14 cards of the deck are off limits
const MaxCardsInPlay = 14

var Rules deck.Rules

const TOTALHANDS = 100_000
const STARTINGBANKROLL = 100000_000
const MINIMUMBET = 100
const SHOES_PER_DECK = 6

func main() {
	rand.Seed(time.Now().Unix())
	Rules = deck.MakeRules(true, true, SHOES_PER_DECK, 1.5, 10)
	dealerWins := 0.0
	playerWins := 0.0
	ties := 0.0
	playerRuins := 0
	playerBankroll := STARTINGBANKROLL
	minBankroll := STARTINGBANKROLL
	maxBankroll := 0

	totalHands := 0.0
	for totalHands < TOTALHANDS {
		shoe := deck.MakeShoe(SHOES_PER_DECK)
		for shoe.GetCardsRemainingBeforeCut() > MaxCardsInPlay {
			trueCount := shoe.TrueCount()
			bet := MINIMUMBET
			if trueCount > 0 {
				if trueCount > Rules.MaxBetSpread() {
					bet *= Rules.MaxBetSpread()
				} else {
					bet *= trueCount //note this doesn't increase the bet for truecount = 1
				}
			}

			//fmt.Println("\n\n\n--- NEW GAME ---\n\n\n")
			if playerBankroll <= 0 {
				playerRuins++
				log.Fatal("PLAYER RUIN")
			}
			allResults := playHand(&shoe, bet)

			playerBankroll += allResults.betOutcome
			if playerBankroll < minBankroll {
				minBankroll = playerBankroll
			} else if playerBankroll > maxBankroll {
				maxBankroll = playerBankroll
			}
			totalHands++

			//if res == DEALERWIN {
			//	dealerWins++
			//} else if res == PLAYERWIN {
			//	playerWins++
			//} else {
			//	ties++
			//}
		}

	}
	log.Printf("Dealer wins %0.2f%%, player wins %0.2f%%, ties: %0.2f%%", 100.0*dealerWins/totalHands,
		100.0*playerWins/totalHands, 100.0*ties/totalHands)

	log.Printf("Final player bankroll: %d (+%0.1f%% of starting)\n", playerBankroll, float64(playerBankroll-STARTINGBANKROLL)/float64(STARTINGBANKROLL)*100.0)
	dollarsPerHand := (float64(playerBankroll) - float64(STARTINGBANKROLL)) / totalHands
	log.Printf("%f dollars increase per hand", dollarsPerHand)
	log.Printf("Minimum bankroll: %d (-%0.1f%% of starting)\n", minBankroll, float64(STARTINGBANKROLL-minBankroll)/float64(STARTINGBANKROLL)*100.0)
	log.Printf("Maximum bankroll: %d (+%0.1f%% of starting)\n", maxBankroll, float64(maxBankroll-STARTINGBANKROLL)/float64(STARTINGBANKROLL)*100.0)

}

type GameResult byte

const (
	PLAYERWIN  GameResult = 'P'
	DEALERWIN  GameResult = 'D'
	PUSH       GameResult = 'T'
	UNFINISHED GameResult = 'U'
)

type HandResults struct {
	hands      []*deck.Hand
	results    []GameResult
	bets       []int
	betOutcome int
}

func playSplitHand(shoe *deck.Shoe, hand *deck.Hand, dealerCard deck.Card, bet int) (handResults HandResults) {
	handResults.hands = make([]*deck.Hand, 1)
	handResults.hands[0] = hand
	handResults.results = make([]GameResult, 1)
	handResults.bets = make([]int, 1)
	handResults.bets[0] = bet

	for !hand.IsBust() {
		//fmt.Println("Player:")
		//fmt.Println(playerHand.ToAscii(false))
		idealMove := deck.PlayerStrategy(Rules, *hand, dealerCard)
		//fmt.Println("Ideal move: " + string(idealMove))
		//action := deck.GetPlayerInput()
		action := idealMove
		if action == deck.HIT {
			card, err := shoe.PullRandomCard()
			if err != nil {
				fmt.Println(err)
			}
			hand.Add(card)
		} else if action == deck.STAND {
			break
		} else if action == deck.DOUBLE {
			handResults.bets[0] *= 2
			card, err := shoe.PullRandomCard()
			if err != nil {
				fmt.Println(err)
			}
			hand.Add(card)
			break
		} else if action == deck.SPLIT {
			log.Println("PLAYING SPLIT HAND")
			hand1, hand2 := hand.Split()
			handResults.hands = make([]*deck.Hand, 0)
			handResults.results = make([]GameResult, 0)
			handResults.bets = make([]int, 0)

			card1Outcome := playSplitHand(shoe, &hand1, dealerCard, bet)
			card2Outcome := playSplitHand(shoe, &hand2, dealerCard, bet)

			handResults.hands = append(handResults.hands, card1Outcome.hands...)
			handResults.results = append(handResults.results, card1Outcome.results...)
			handResults.bets = append(handResults.bets, card1Outcome.bets...)
			handResults.betOutcome += card1Outcome.betOutcome

			handResults.hands = append(handResults.hands, card2Outcome.hands...)
			handResults.results = append(handResults.results, card2Outcome.results...)
			handResults.bets = append(handResults.bets, card2Outcome.bets...)
			handResults.betOutcome += card2Outcome.betOutcome
			break
		} else {
			break
		}
	}
	return handResults
}

func playHand(shoe *deck.Shoe, bet int) (allResults HandResults) {
	allResults.hands = make([]*deck.Hand, 1)
	allResults.results = make([]GameResult, 1)
	allResults.betOutcome = 0
	allResults.bets = make([]int, 1)
	allResults.bets[0] = bet

	//burn card
	_, err := shoe.PullRandomCard()
	if err != nil {
		log.Fatal(err)
	}

	player1, err := shoe.PullRandomCard()

	playerHand := deck.MakeHand(player1)
	dealer := deck.MakeDealerAndHand(Rules.DoesDealerHitSoft17(), shoe)
	player2, err := shoe.PullRandomCard()
	playerHand.Add(player2)
	dealerDown, err := shoe.PullRandomCard()
	dealer.AddDownCard(dealerDown)
	allResults.hands[0] = &playerHand

	//TODO Implement insurance?
	if playerHand.HasBlackjack() && dealer.Hand().HasBlackjack() {
		log.Println("Both have blackjack")
		allResults.betOutcome += 0
		allResults.results[0] = PUSH
		return allResults
	} else if dealer.Hand().HasBlackjack() {
		log.Println("Dealer alone has blackjack")
		allResults.betOutcome += -bet
		allResults.results[0] = DEALERWIN
		return allResults
	} else if playerHand.HasBlackjack() {
		log.Println("Player has blackjack")
		allResults.betOutcome += int(float64(bet) * Rules.BlackjackPayout())
		allResults.results[0] = PLAYERWIN
		return allResults
	}

	results := playSplitHand(shoe, &playerHand, dealer.UpCard(), bet)

	//fmt.Println("Dealer:")
	//fmt.Println(dealer.UpCard().ToAscii() + "  ?")

	for dealer.DoesHit() {
		dealer.Hit(shoe)
	}
	dealerHand := dealer.Hand()

	allResults.results = make([]GameResult, len(results.hands))
	for i, v := range results.hands {
		if len(results.hands) > 1 {
			log.Println(v.ToString(false))
		}
		if v.IsBust() || (dealerHand.HighestPlay() > v.HighestPlay() && !v.IsBust()) {
			allResults.betOutcome += -results.bets[i]
			allResults.results[i] = DEALERWIN
		} else if dealerHand.HighestPlay() < v.HighestPlay() || v.IsBust() {
			allResults.betOutcome += results.bets[i]
			allResults.results[i] = PLAYERWIN
		} else {
			allResults.betOutcome += 0
			allResults.results[i] = PUSH
		}
	}

	return allResults
}
