package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
}

func main() {
	demoSwitch()
	demoSelect()
}

func demoSwitch() {
	var a interface{}
	a = 5

	switch a.(type) {
	case int:
		fmt.Println("It's int")
	case Person:
		fmt.Println("It's a person")
	default:
		fmt.Println("Not matching any!")
	}
}

func demoSelect() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	// infinite for loop
	for {

		// select is used with reading from channels
		// it executes which ever channels gets the value first
		// it goes to default if it's defined when no other channel gives a value
		select {
		case <-tick:
			fmt.Println("Tick.")
		case <-boom:
			fmt.Println("Boom!")
			return
		default:
			time.Sleep(50 * time.Millisecond)
			fmt.Println("   .")
		}
	}
}
