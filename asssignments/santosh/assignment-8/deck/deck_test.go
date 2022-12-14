package deck

import (
	"log"
	"testing"
)

// unit tests for creating a deck

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

	// case 3: the last card should be Ace of Spades
	if newDeck[15] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs but got %v", newDeck[0])
	}

	// case 4: there should be Two of hearts in the deck
	var foundAt int
	for ind, name := range newDeck {
		if name == "Two of Hearts" {
			foundAt = ind
		}
	}

	if foundAt == 0 {
		t.Errorf("Expected Two of Hearts but not found in the deck ")
	}

}

// Unit test of Shuffle function

//case 1: deck size should remain same
//case 2: all the elements should remain same
//case 3: pass the deck and except shuffled deck

func TestShuffle(t *testing.T) {
	originaDeck := CreateNewDeck()
	shuffledDeck := make([]string, len(originaDeck))
	copy(shuffledDeck, originaDeck)
	shuffledDeck = Shuffle(shuffledDeck)

	// case 1
	if len(originaDeck) != len(shuffledDeck) {
		t.Errorf("Expected size is 16 but got %d", len(shuffledDeck))
	}

	//case 2
	for _, originalCard := range originaDeck {
		if !contains(shuffledDeck, originalCard) {
			t.Errorf("Expected same elements in the shuffled deck but found different")
		}
	}

	//case 3
	var count int
	for i := 0; i < 16; i++ {
		if originaDeck[i] == shuffledDeck[i] {
			log.Println(originaDeck[i], shuffledDeck[i])
			count = count + 1
		}
	}

	if count == 16 {
		t.Errorf("Excepted shuffled cards but found same ordered deck %d", count)
	}
}

// function to indentify check a card present in deck

func contains(deck []string, card string) bool {
	var status bool
	for _, deckCard := range deck {
		if card == deckCard {
			status = true
		}
	}
	return status
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

// Reading file

// case 1: empty file
// case 2: wrong file name
func TestReadFile(t *testing.T) {
	err := ReadFile("")
	if err == nil {
		t.Error("Expected error: filename cannot be empty")
	}

	err1 := ReadFile("abc.txt")
	if err1 == nil {
		t.Error("Expected error: filename cannot be empty")
	}

}

// deal function
func TestDeal(t *testing.T) {
	deck := CreateNewDeck()
	Deal(deck, -2)
	Deal(deck, 16)
}
