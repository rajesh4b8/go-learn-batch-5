package main

import "fmt"

func main() {
	a := 23

	if a%2 == 0 {
		fmt.Println("It's even")
	} else {
		fmt.Println("It's odd")
	}

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
