package deck

import (
	"errors"
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
	return newDeck
}

// using rand package to generate random index
func Shuffle(deck Deck) Deck {
	log.Println("recived deck", deck)
	rand.Seed(time.Now().UnixNano())
	for index := range deck {
		randomIndex := rand.Intn(16)
		//log.Println(index, randomIndex)
		deck[index], deck[randomIndex] = deck[randomIndex], deck[index]
	}
	log.Println("returned deck", deck)
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

func Deal(d Deck, handSize int) (Deck, Deck, error) {

	if handSize >= 16 {
		//log.Println("Handsize cannot be more than 16")
		return d, d, errors.New("Handsize cannot be more than 16")
	}
	if handSize <= 0 {
		//log.Println("Handsize cannot be less than or equal to 0")
		return d, d, errors.New("Handsize cannot be less than or equal to 0")
	}
	return d[:handSize], d[handSize:], nil
}

func (d Deck) WriteToFile(fileName string) error {
	if len(fileName) == 0 {
		return errors.New("fileName can't be empty")
	} else if !strings.Contains(fileName, ".") {
		return errors.New("fileName should be with a . for proper extn")
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

func ReadFile(fileName string) (d Deck, err error) {
	log.Println(fileName)
	if len(fileName) == 0 {
		//log.Println("Error in reading file: File name cannot be empty")
		var emptyDeck Deck
		return emptyDeck, errors.New("Error in reading file: File name cannot be empty")
	} else if !strings.Contains(fileName, ".") {
		//log.Println("Error in reading file: file name should have extention ")
		var emptyDeck Deck
		return emptyDeck, errors.New("Error in reading file: file name should have extention")
	}

	content, err := ioutil.ReadFile(fileName)
	log.Println("err", err)
	//fmt.Println("err", err)
	if err != nil {
		//log.Println("Error in reading file : ", err)
		var emptyDeck Deck
		return emptyDeck, err
	}

	//fmt.Println("File content is ", string(content))
	deck := strings.Split(string(content), ",")
	//fmt.Println(deck, len(deck), cap(deck))
	return deck, err

}

// func main() {
// 	fmt.Println(CreateNewDeck())
// }
