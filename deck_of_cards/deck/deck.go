package deck

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Deck []string

func CreateNewDeck() Deck {
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	numbers := []string{"Ace", "Two", "Three", "Four"}

	// created with size 0
	newDeck := Deck{}

	// This inializes with zero values
	// newDeck := make([]string, 16)

	// fmt.Println("newDeck", "Length:", len(newDeck), "Cap:", cap(newDeck), "Contents:", newDeck)

	for _, suit := range suits {
		for _, num := range numbers {
			newDeck = append(newDeck, num+" of "+suit)
		}
	}
	// fmt.Println("slice1", "Length:", len(newDeck), "Cap:", cap(newDeck), "Contents:", newDeck)

	return newDeck
}

func (deck Deck) Shuffle() {
	// TODO: logic for shuffle

	// how to switch values?
	// deck[0], deck[4] = deck[4], deck[0]
}

// Method on Deck
func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) WriteToFile(fileName string) error {
	if len(fileName) == 0 {
		return errors.New("fileName can't be empty")
	} else if !strings.Contains(fileName, ".") {
		return errors.New("fileName should be with a `.` for proper extn")
	}

	str := strings.Join(d, ",")
	content := []byte(str)

	err := ioutil.WriteFile(fileName, content, 0666)
	if err != nil {
		log.Fatal("Error writiting to a file", err)
		os.Exit(1)
	}

	return nil
}

func ReadFromFile(fileName string) Deck {
	// TODO:
	// Read file using the ioutil package
	// Convert []byte to string
	// split string to []string

	return nil
}
