package main

import (
	"fmt"

	"github.com/rajesh4b8/go-learn-batch-5/asssignments/santosh/assignment-5/deck"
	dealPkg "github.com/rajesh4b8/go-learn-batch-5/asssignments/santosh/assignment-5/deck"
)

func main() {
	newDeck := deck.CreateNewDeck()
	//newDeck.Print()

	//deck.Print(newDeck)
	//deck.Print(deck.Shuffle(newDeck))
	// fmt.Println(newDeck)
	// fmt.Println(deck.Shuffle(newDeck))

	deal, remaining := dealPkg.Deal(newDeck, 5)
	fmt.Println("Deal cards are:")
	deal.Print()
	fmt.Println("Remaining cards are:")
	remaining.Print()
	// To write deal card data in to the text file
	deal.WriteToFile("DealCards.txt")

	// To read deal card data from the text file
	fmt.Println("******** Reading Dealcard data from file ********")
	dealDeckFromFile := dealPkg.ReadFile("DealCards.txt")
	dealDeckFromFile.Print()

}
