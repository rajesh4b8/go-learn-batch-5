package main

import (
	"fmt"
)

var (
	a   int = 0
	b   int = 1
	n   int = 6
	num int = 0
)

func main() {
	if n < 1 {
		fmt.Println("Enter a number greater than 1")
	} else if n == 1 {
		fmt.Println(a)
	} else if n == 2 {
		fmt.Print(a, " ", b)
	} else {
		fmt.Print(a, " ", b, " ")
		for i := 3; i <= n; i++ {
			num = a + b
			fmt.Print(num, " ")
			a = b
			b = num
		}
	}
}
