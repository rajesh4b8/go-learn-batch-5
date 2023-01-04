package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// go is the keyword here to make the func running in a child goroutine
	go say("Go Hello")

	// Either you can use Channels or WaitGroups to hold the main goroutine

	say(" Main Hello")

}
