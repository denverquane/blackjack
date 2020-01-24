package main

import (
	"github.com/denverquane/blackjack/deck"
	"log"
	"math/rand"
	"time"
)

const SEED = 1234567

//Dumb heuristic, assume the last 14 cards of the deck are off limits
const MaxCardsInPlay = 30

var Rules deck.Rules

const TOTALHANDS = 1000000
const STARTINGBANKROLL = 10000000000
const MINIMUM_BET = 100
const SHOES_PER_DECK = 6

func main() {
	rand.Seed(time.Now().Unix())
	Rules = deck.MakeRules(true, false, SHOES_PER_DECK, 1.5, 100, false)
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
			trueCount := shoe.TrueCount() - 1
			bet := MINIMUM_BET
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
			allResults, betSum := playHand(&shoe, bet)

			playerBankroll += betSum

			if playerBankroll < minBankroll {
				minBankroll = playerBankroll
			} else if playerBankroll > maxBankroll {
				maxBankroll = playerBankroll
			}
			totalHands += float64(allResults[PLAYERWIN])
			playerWins += float64(allResults[PLAYERWIN])

			totalHands += float64(allResults[DEALERWIN])
			dealerWins += float64(allResults[DEALERWIN])

			totalHands += float64(allResults[PUSH])
			ties += float64(allResults[PUSH])

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

type AllResults map[GameResult]int

const (
	PLAYERWIN GameResult = 'P'
	DEALERWIN GameResult = 'D'
	PUSH      GameResult = 'T'
)

func playHand(shoe *deck.Shoe, bet int) (results AllResults, betSum int) {
	//burn card
	_, err := shoe.PullRandomCard()
	if err != nil {
		log.Fatal(err)
	}

	player1, err := shoe.PullRandomCard()

	playerHand := deck.MakeHand(player1)

	dealer := deck.MakeDealerAndHand(Rules.DoesDealerHitSoft17(), shoe)

	shoe.PullAndAddToHand(playerHand)

	dealerDown, err := shoe.PullRandomCard()
	dealer.AddDownCard(dealerDown)

	//TODO Implement insurance?

	results, betSum = playOutEntireHand(shoe, playerHand, dealer, bet, 0)

	return results, betSum
}

func playOutEntireHand(shoe *deck.Shoe, playerHand *deck.Hand, dealer deck.Dealer, bet, recursionLevel int) (results AllResults, betSum int) {
	results = make(AllResults)
	results[PLAYERWIN] = 0
	results[DEALERWIN] = 0
	results[PUSH] = 0

	if recursionLevel == 0 {
		if playerHand.HasBlackjack() && dealer.Hand().HasBlackjack() {
			log.Println("Both have blackjack")
			results[PUSH]++
			return results, 0
		} else if dealer.Hand().HasBlackjack() {
			log.Println("Dealer alone has blackjack")
			results[DEALERWIN]++
			return results, -bet
		} else if playerHand.HasBlackjack() {
			log.Println("Player has blackjack")
			results[PLAYERWIN]++
			return results, int(float64(bet) * Rules.BlackjackPayout())
		}
	}

	for !playerHand.IsBust() {

		action := deck.PlayerStrategy(Rules, *playerHand, dealer.Hand().FirstCard())

		if action == deck.STAND {
			break
			//if hitting, OR the hand is split, but resplits aren't allowed
		} else if action == deck.HIT || (action == deck.SPLIT && !Rules.Resplit() && playerHand.IsSplit()) {
			shoe.PullAndAddToHand(playerHand)

			//don't exit, can keep hitting
		} else if action == deck.DOUBLE {
			bet *= 2.0
			shoe.PullAndAddToHand(playerHand)

			break
		} else if action == deck.SPLIT {
			log.Printf("PLAYING SPLIT HAND level %d\n", recursionLevel)
			hand1, hand2 := playerHand.Split()

			shoe.PullAndAddToHand(hand1)
			shoe.PullAndAddToHand(hand2)

			log.Println(hand1.ToAscii(false))
			log.Println(hand2.ToAscii(false))

			split1Result, split1Bet := playOutEntireHand(shoe, hand1, dealer, bet, recursionLevel+1)
			split2Result, split2Bet := playOutEntireHand(shoe, hand2, dealer, bet, recursionLevel+1)

			bet = split1Bet + split2Bet

			for i, v := range split1Result {
				results[i] += v
			}
			for i, v := range split2Result {
				results[i] += v
			}

			return results, bet
		} else {
			break
		}
	}
	for dealer.DoesHit() {
		dealer.Hit(shoe)
	}

	if playerHand.IsBust() || (dealer.Hand().HighestPlay() > playerHand.HighestPlay()) {
		results[DEALERWIN]++
		return results, -bet
	} else if dealer.Hand().HighestPlay() < playerHand.HighestPlay() || dealer.Hand().IsBust() {
		results[PLAYERWIN]++
		return results, bet
	} else {
		results[PUSH]++
		return results, 0
	}
}
