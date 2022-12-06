package main

import "fmt"

type card string

func main() {
	c := card("Ace of Spaes")
	printMeFunc(c)
	c.printMeMethod()
}

// this is a function and no receiver, called with package name
func printMeFunc(c card) {
	fmt.Println("Card is", c)
}

// A Reciever varable c is defined on method
func (c card) printMeMethod() {
	fmt.Println("Card is", c)
}
