package main

import "fmt"

func main() {

	var factorialnum, factorial int
	factorial = 1

	fmt.Print("Enter any Number to find the Factorial = ")
	fmt.Scanln(&factorialnum)

	for i := 1; i <= factorialnum; i++ {
		factorial = factorial * i
	}
	fmt.Println("The Factorial of ", factorialnum, " = ", factorial)
}
