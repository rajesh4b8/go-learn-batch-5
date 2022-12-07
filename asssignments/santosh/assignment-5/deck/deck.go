package deck

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Deck []string

func CreateNewDeck() Deck {
	suits := Deck{"Spades", "Hearts", "Diamonds", "Clubs"}
	numbers := Deck{"Ace", "Two", "Three", "Four"}

	// created with size 0
	//newDeck := Deck{}
	newDeck := make(Deck, 16)

	for i, suit := range suits {
		for j, num := range numbers {
			//newDeck = append(newDeck, num+" of "+suit)
			// fmt.Println(i, j, i*4+j, "\t", "inside loop")
			newDeck[i*4+j] = num + " of " + suit
		}
	}
	return Shuffle(newDeck)
}

// using rand package to generate random index
func Shuffle(deck Deck) Deck {
	rand.Seed(time.Now().UnixNano())
	for index := range deck {
		randomIndex := rand.Intn(index + 1)
		deck[index], deck[randomIndex] = deck[randomIndex], deck[index]
	}
	return deck
}

func Print(d Deck) {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) WriteToFile(fileName string) {
	//file_byte = byte[](d)
	file_data := []byte(strings.Join(d, ","))
	err := ioutil.WriteFile(fileName, file_data, 0666)
	if err != nil {
		log.Fatal("Error writing to file", err)
		os.Exit(1)
	}
	fmt.Println("Dealcard data has been written to file successfully")
}

func ReadFile(fileName string) (d Deck) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error in reading file ", err)
		os.Exit(1)
	}

	//fmt.Println("File content is ", string(content))
	deck := strings.Split(string(content), ",")
	//fmt.Println(deck, len(deck), cap(deck))
	return deck

}

// func main() {
// 	fmt.Println(CreateNewDeck())
// }
