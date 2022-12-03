package deck

// type Deck []string

func CreateNewDeck() []string {
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	numbers := []string{"Ace", "Two", "Three", "Four"}

	newDeck := []string{}

	for _, suit := range suits {
		for _, num := range numbers {
			newDeck = append(newDeck, num+" of "+suit)
		}
	}

	return newDeck
}

func Shuffle(deck []string) []string {
	// TODO: logic for shuffle

	// how to switch values?
	// deck[0], deck[4] = deck[4], deck[0]
	return deck
}
