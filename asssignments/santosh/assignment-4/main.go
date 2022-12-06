package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateNewDeck() []string {
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	numbers := []string{"Ace", "Two", "Three", "Four"}

	// created with size 0
	newDeck := []string{}

	for _, suit := range suits {
		for _, num := range numbers {
			newDeck = append(newDeck, num+" of "+suit)
		}
	}
	return Shuffle(newDeck)
}

func Shuffle(deck []string) []string {
	rand.Seed(time.Now().UnixNano())
	for index := range deck {
		randomIndex := rand.Intn(index + 1)
		deck[index], deck[randomIndex] = deck[randomIndex], deck[index]
	}
	return deck
}

func main() {
	fmt.Println(CreateNewDeck())
}
