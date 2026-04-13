package main

import (
	"fmt"

	"github.com/Gustavo-DCosta/tarot.go/models"
	"github.com/Gustavo-DCosta/tarot.go/services"
)

func main() {
	Cards := models.NewDeck()
	// freshCards stands for the name of the array of cards that has been cut
	// I couldn't find a better name and it was 22:26
	freshCards := services.CutCards(Cards)
	shufledCards, players, dog := services.DealCards(freshCards)

	/*for _, card := range shufledCards {
		// This prints each card on its own line
		fmt.Printf("%+v\n", card)
	}*/
	_ = shufledCards

	fmt.Print()

	for _, plcard := range players {
		// This prints each card on its own line
		fmt.Printf("%+v\n", plcard)
	}

	for _, dcard := range dog {
		// This prints each card on its own line
		fmt.Printf("%+v\n", dcard)
	}

}
