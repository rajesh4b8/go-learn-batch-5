package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateNewDeck() []string {
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	numbers := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// created with size 0
	newDeck := []string{}

	for _, suit := range suits {
		for _, num := range numbers {
			newDeck = append(newDeck, num+" of "+suit)
		}
	}
	return Shuffle(newDeck)
}

// using rand package to generate random position
func Shuffle(deck []string) []string {
	rand.Seed(time.Now().UnixNano())
	for pos := range deck {
		randomPos := rand.Intn(len(deck))
		deck[pos], deck[randomPos] = deck[randomPos], deck[pos]
	}
	return deck
}

// Main Function.
func main() {
	fmt.Println(CreateNewDeck()) //Function calling.
}
