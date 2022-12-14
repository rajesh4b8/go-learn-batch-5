package deck

import (
	"log"
	"testing"
)

// TDD -> Test Driven Development???
// Write your tests first
// know what to be expected out of your code

// Unit tests for a multiplication
// func multiply(x, y int) int
// use case 1: 2, 3 should return 6
// use case 2: -2, 3 should return -6

// unit tests for creating a deck

// case 3: the last card should be Four of Clubs

func TestCreateNewDeck(t *testing.T) {
	newDeck := CreateNewDeck()

	// case 1: there should be 16 cards
	if len(newDeck) != 16 {
		t.Errorf("Expected size is 16 but got %d", len(newDeck))
	}

	// case 2: the first card should be Ace of Spades
	if newDeck[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", newDeck[0])
	}
}

// unit test for writing to a file

// Case 1: pass file name get with the deck and expect a file to be created
// case 2: return error when file name is empty

func TestWriteToFile(t *testing.T) {

	deck := CreateNewDeck()

	// case 2: return error when file name is empty
	err := deck.WriteToFile("")
	log.Fatal(err)

	// it's supposed return error when the input is empty which is not valid
	if err == nil {
		t.Error("Epecting an error here but never returned")
	}
}
