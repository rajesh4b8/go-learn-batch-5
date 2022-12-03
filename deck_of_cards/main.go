package main

import (
	"fmt"

	"github.com/rajesh4b8/go-learn-batch-5/deck_of_cards/deck"
)

func main() {

	newDeck := deck.CreateNewDeck()
	fmt.Println(newDeck)
	fmt.Println(deck.Shuffle(newDeck))
}
