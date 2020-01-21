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

	shoe := deck.MakeShoe(2)

	for {
		for shoe.GetCardsRemainingBeforeCut() > MaxCardsInPlay {
			fmt.Println("\n\n\n--- NEW GAME ---\n\n\n")
			playHand(shoe)
		}
		shoe = deck.MakeShoe(2)
	}
}

func playHand(shoe deck.Shoe) {
	//burn card
	_, err := shoe.PullRandomCard()
	if err != nil {
		log.Fatal(err)
	}

	player1, err := shoe.PullRandomCard()

	playerHand := deck.MakeHand(player1)
	dealer := deck.MakeDealerAndHand(DealerHitsSoft17, &shoe)
	player2, err := shoe.PullRandomCard()
	playerHand.Add(player2)

	fmt.Println("Dealer:")
	fmt.Println(dealer.UpCard().ToAscii() + "  ?")

	for !playerHand.IsBust() {
		fmt.Println("Player:")
		fmt.Println(playerHand.ToAscii(false))
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
	fmt.Println(playerHand.ToAscii(false))

	for dealer.DoesHit(&shoe) {

	}
	fmt.Println("Dealer:")
	dealerHand := dealer.Hand()
	fmt.Println(dealerHand.ToAscii(false))

	if playerHand.IsBust() || (dealerHand.HighestPlay() > playerHand.HighestPlay() && !dealerHand.IsBust()) {
		fmt.Println("--- DEALER WINS ---")
	} else if dealerHand.HighestPlay() < playerHand.HighestPlay() || dealerHand.IsBust() {
		fmt.Println("--- PLAYER WINS ---")
	} else {
		fmt.Println("PUSH")
	}
}
