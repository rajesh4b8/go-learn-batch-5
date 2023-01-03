package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	start = time.Now()
)

func main() {
	links := []string{
		"http://linkedin.com",
		"http://google.com",
		"http://facebook.com",
		"http://learningnet.com",
		"http://yahoo.com",
	}

	// chan is the keyword
	// creating a channel of type string by using make()
	c := make(chan string) // non-buffered channel
	for _, l := range links {
		// spin a new child goroutine to check the health
		go checkHealth(l, c)
	}

	for range links {
		fmt.Println(<-c)
	}
}

func checkHealth(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- fmt.Sprintf("%s Might be down! - in %v\n", link, time.Since(start))
		return
	}

	c <- fmt.Sprintf("%s is up! - in %v\n", link, time.Since(start))
}
