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
const STARTINGBANKROLL = 10_000
const MINIMUMBET = 100
const SHOES_PER_DECK = 8

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
	for totalHands < TOTALHANDS{
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
			res, betResult := playHand(&shoe, bet)
			playerBankroll += betResult
			if playerBankroll < minBankroll {
				minBankroll = playerBankroll
			} else if playerBankroll > maxBankroll {
				maxBankroll = playerBankroll
			}
			totalHands ++
			if res == DEALERWIN {
				dealerWins++
			} else if res == PLAYERWIN {
				playerWins++
			} else {
				ties++
			}
		}

	}
	log.Printf("Dealer wins %0.2f%%, player wins %0.2f%%, ties: %0.2f%%", 100.0*dealerWins / totalHands,
		100.0*playerWins / totalHands, 100.0*ties / totalHands)

	log.Printf("Final player bankroll: %d (+%0.1f%% of starting)\n", playerBankroll, float64(playerBankroll - STARTINGBANKROLL) / float64(STARTINGBANKROLL) * 100.0)
	dollarsPerHand :=  (float64(playerBankroll) - float64(STARTINGBANKROLL)) / totalHands
	log.Printf("%f dollars increase per hand", dollarsPerHand)
	log.Printf("Minimum bankroll: %d (-%0.1f%% of starting)\n", minBankroll, float64(STARTINGBANKROLL - minBankroll) / float64(STARTINGBANKROLL) * 100.0)
	log.Printf("Maximum bankroll: %d (+%0.1f%% of starting)\n", maxBankroll, float64(maxBankroll - STARTINGBANKROLL) / float64(STARTINGBANKROLL) * 100.0)

}

type GameResult byte

const (
	PLAYERWIN GameResult = 'P'
	DEALERWIN GameResult = 'D'
	PUSH GameResult = 'T'
)

func playHand(shoe *deck.Shoe, bet int) (res GameResult, betResult int) {

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


	//TODO Implement insurance?
	if playerHand.HasBlackjack() && dealer.Hand().HasBlackjack() {
		log.Println("Both have blackjack")
		return PUSH, 0
	} else if dealer.Hand().HasBlackjack() {
		log.Println("Dealer alone has blackjack")
		return DEALERWIN, -bet
	} else if playerHand.HasBlackjack() {
		log.Println("Player has blackjack")
		return PLAYERWIN, int(float64(bet) * Rules.BlackjackPayout())
	}

	//fmt.Println("Dealer:")
	//fmt.Println(dealer.UpCard().ToAscii() + "  ?")

	for !playerHand.IsBust() {
		//fmt.Println("Player:")
		//fmt.Println(playerHand.ToAscii(false))
		idealMove := deck.PlayerStrategy(Rules, playerHand, dealer.UpCard())
		//fmt.Println("Ideal move: " + string(idealMove))
		//action := deck.GetPlayerInput()
		action := idealMove
		if action == deck.HIT {
			card, err := shoe.PullRandomCard()
			if err != nil {
				fmt.Println(err)
			}
			playerHand.Add(card)
		} else if action == deck.STAND {
			break
		} else if action == deck.DOUBLE {
			bet *= 2
			card, err := shoe.PullRandomCard()
			if err != nil {
				fmt.Println(err)
			}
			playerHand.Add(card)
			break
			//TODO Implement Splits
		} else if action == deck.SPLIT && playerHand.HighestPlay() < 17 {
			card, err := shoe.PullRandomCard()
			if err != nil {
				fmt.Println(err)
			}
			playerHand.Add(card)
		} else {
			break
		}
	}

	for dealer.DoesHit() {
		dealer.Hit(shoe)
	}
	dealerHand := dealer.Hand()

	if playerHand.IsBust() || (dealerHand.HighestPlay() > playerHand.HighestPlay() && !dealerHand.IsBust()) {
		return DEALERWIN, -bet
	} else if dealerHand.HighestPlay() < playerHand.HighestPlay() || dealerHand.IsBust() {
		return PLAYERWIN, bet
	} else {
		return PUSH, 0
	}
}
