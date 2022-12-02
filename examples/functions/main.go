package main

import "fmt"

func main() {
	// calling the function with `()`
	simpleFunction()
	withParams("Rajesh", 12)
	withParamsSameType(2, 4)
	fmt.Println("Returned string:", retunAString())
	x := add(20, 40)
	fmt.Println("20 + 40 =", x)

	name, age := returningMultiValues()
	fmt.Println(name, age)
}

// func is the keyword
// simpleFunction is the name of the function
// this is the defintion of simpleFunction
func simpleFunction() {
	fmt.Println("Simple function")
}

// function with input parameters
func withParams(name string, age int) {
	fmt.Println("Details:", name, age)
}

func withParamsSameType(i, j int) {
	fmt.Println("Values:", i, j)
}

func retunAString() string {
	return "Some string"
}

func add(i, j int) int {
	return i + j
}

func returningMultiValues() (string, int) {
	return "Rajesh", 12
}
