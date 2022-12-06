package main

import (
	"fmt"

	dealPkg "github.com/rajesh4b8/go-learn-batch-5/deck_of_cards/deck"
)

func main() {

	newDeck := dealPkg.CreateNewDeck()
	// fmt.Println(deck.Shuffle())

	newDeck.Print()

	deal, remaining := dealPkg.Deal(newDeck, 5)
	fmt.Println("Deal cards", deal)
	fmt.Println("Remaining cards", remaining)

	deal.WriteToFile("DealCards.txt")
}
