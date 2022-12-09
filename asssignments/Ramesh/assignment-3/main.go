package main

import "fmt"

func main() {
	printFibSeries(15)

}

// Print the n number of elements in fib series
func printFibSeries(n int) {
	StartNum, secondNum := 1, 2

	if n < 1 {
		fmt.Println("Invalid number of elements")
	} else if n == 1 {
		fmt.Println("Element 1 is ", StartNum)
	} else if n == 2 {
		fmt.Println("Element 1 is ", StartNum)
		fmt.Println("Element 2 is ", secondNum)
	} else {
		fmt.Println("Element 1 is", StartNum)
		fmt.Println("Element 2 is", secondNum)
		for i := 3; i <= n; i++ {
			result := StartNum + secondNum
			fmt.Printf("Element %d is %d\n", i, result)
			StartNum = secondNum
			secondNum = result
		}
	}
}
