package services

import (
	"fmt"
	"math/rand/v2"

	"github.com/Gustavo-DCosta/tarot.go/models"
	"github.com/fatih/color"
)

/*
	In Tarot we distribute n cards for each time to each player
	First 3 cards are giving to the 5 players then one goes to the dog (chien)
	We enter in a loop where we feed the players 3 cards each and the dog 1 card for 3 cycles
	Then we go back to feed the players hand 3 cards each time until everyone has 15
*/

func DealCards(cards []models.Card) ([]models.Card, [][]models.Card, []models.Card) {
	// 1. Shuffle first
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	// 2. Initialize players as a slice of slices (easier to loop through)
	players := make([][]models.Card, 5)
	var dog []models.Card

	currentCard := 0

	// 3. Deal 3 cards to each player at a time (Standard for games like Tarot)
	// We do this until everyone has their hand (adjust the loop for your rules)
	for round := 0; round < 5; round++ { // Example: 6 rounds of 3 cards = 18 cards each
		for p := 0; p < 5; p++ {
			for k := 0; k < 3; k++ {
				players[p] = append(players[p], cards[currentCard])
				currentCard++
			}
		}
	}

	// 4. Deal to the "dog"
	for d := 0; d < 3; d++ {
		dog = append(dog, cards[currentCard])
		currentCard++
	}

	// Return the players and the dog (and maybe the remaining deck)
	return cards, players, dog
}

func CutCards(cards []models.Card) []models.Card {
	x := 10
	n := len(cards) // Better to use the actual slice length!

	actionColor := color.RGB(110, 175, 114)
	actionColor.Println("...Cutting the deck...")

	// 1. Get your random pivot point
	randomNum := rand.IntN(n-x) + x
	fmt.Printf("Cutting at: %d\n", randomNum)

	// 2. The "Go Way": Re-slice and concatenate
	// We take everything FROM randomNum to the end [randomNum:]
	// And append everything from the START to randomNum [:randomNum]
	cutDeck := append(cards[randomNum:], cards[:randomNum]...)
	// Just for debug

	/*for _, card := range cutDeck {
		// This prints each card on its own line
		fmt.Printf("%+v\n", card)
	}*/

	return cutDeck
}
