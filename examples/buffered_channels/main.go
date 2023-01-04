package main

import "fmt"

func main() {
	// non-buffered channel
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	fmt.Println("Non buffered channel", <-ch)

	bc := make(chan int, 4)
	bc <- 1
	bc <- 2
	bc <- 3

	fmt.Println("buffered channel", <-bc)

	bc <- 4
	bc <- 5

	ch1 := make(chan int)

	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		close(ch1)
	}()

	for i := range ch1 {
		fmt.Println(i)
	}

}
