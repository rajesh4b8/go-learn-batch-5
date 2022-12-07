package main

import "fmt"

func main() {
	printFibSeries(5)
}

// Print the n number of elements in fib seiries
func printFibSeries(n int) {
	firstNum, secNum := 0, 1

	if n < 1 {
		fmt.Println("Invalid number of elements")
	} else if n == 1 {
		fmt.Println("Element 1 is ", firstNum)
	} else if n == 2 {
		fmt.Println("Element 1 is ", firstNum)
		fmt.Println("Element 2 is ", secNum)
	} else {
		fmt.Println("Element 1 is", firstNum)
		fmt.Println("Element 2 is", secNum)
		for i := 3; i <= n; i++ {
			res := firstNum + secNum
			fmt.Printf("Element %d is %d\n", i, res)
			firstNum = secNum
			secNum = res
		}
	}
}
