package main

import (
	"fmt"
	"github.com/denverquane/blackjack/deck"
	"log"
	"math/rand"
	"time"
)

const SEED = 1234567

const DealerHitsSoft17 = true

//Dumb heuristic, assume the last 14 cards of the deck are off limits
const MaxCardsInPlay = 14

func main() {
	rand.Seed(time.Now().Unix())
	//
	//singleDeck := deck.MakeDeck()
	//
	//card1, _ := singleDeck.PullRandomCard()
	//card2, _ := singleDeck.PullRandomCard()
	//card3, _ := singleDeck.PullRandomCard()
	//card4, _ := singleDeck.PullRandomCard()
	//hand := deck.MakeHand(card1, card2)
	//hand.Add(card3)
	//hand.Add(card4)
	//fmt.Println(hand.ToAscii())

	shoe := deck.MakeShoe(2)
	if shoe.GetCardsRemainingBeforeCut() > MaxCardsInPlay {
		playHand(shoe)
	}

}

func playHand(shoe deck.Shoe) {
	//burn card
	_, err := shoe.PullRandomCard()
	if err != nil {
		log.Fatal(err)
	}

	dealer1, err := shoe.PullRandomCard()
	player1, err := shoe.PullRandomCard()
	dealer2, err := shoe.PullRandomCard()
	player2, err := shoe.PullRandomCard()

	dealerHand := deck.MakeHand(dealer1, dealer2)
	playerHand := deck.MakeHand(player1, player2)
	dealer := deck.MakeDealer(DealerHitsSoft17)

	fmt.Println("Dealer:")
	fmt.Println(dealerHand.FirstCard().ToAscii() + "  ?")

	for !playerHand.IsBust() {
		fmt.Println("Player:")
		fmt.Println(playerHand.ToAscii())
		action := deck.GetPlayerInput()
		if action == deck.HIT {
			card, err := shoe.PullRandomCard()
			if err != nil {
				fmt.Println(err)
			}
			playerHand.Add(card)
		} else if action == deck.STAND {
			break
		}
	}
	fmt.Println("- Final Player Hand -")
	fmt.Println(playerHand.ToAscii())

	for dealer.DoesHit(*dealerHand) {
		card, err := shoe.PullRandomCard()
		if err != nil {
			fmt.Println(err)
		}
		dealerHand.Add(card)
	}
	fmt.Println("Dealer:")
	fmt.Println(dealerHand.ToAscii())

	if dealerHand.HighestPlay() > playerHand.HighestPlay() {
		fmt.Println("DEALER WINS")
	} else if dealerHand.HighestPlay() < playerHand.HighestPlay() {
		fmt.Println("PLAYER WINS")
	} else {
		fmt.Println("PUSH")
	}
}
