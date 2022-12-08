package main

import "fmt"

// fib series using recursion
func fib_series(n int) int {

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib_series(n-1) + fib_series((n - 2))
	}

}

// reading fibonacci series postion from the console
func main() {
	fmt.Print("Enter the fiibonacci series position ")
	var number int
	fmt.Scanln(&number)
	fmt.Println(fib_series(number))

}
