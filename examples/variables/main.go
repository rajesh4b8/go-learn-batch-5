package main

import "fmt"

var (
	t = 12
)

func main() {
	// var is the keyword
	// foo is the name of the variable
	// int is the data type (built in type)
	// 56 is the value we are initialzing with

	var foo int = 56 // declration with initialization
	// foo = "Test" // Changing the type is not allowed after declaring with a specific type
	var foo1 = 56 // declration with initialization

	var bar int // declration without initialization

	// var i, j int = 2, 3 // multi value declartion with initialization
	var i, j = 2, 3 // multi value declartion with initialization
	var x, y int

	// var name string = "Rajesh"
	name := "Rajesh" // declration with initialization
	name = "Ramesh"  // assignment / change value

	// var age int // declartion
	// age = 20 // assignment
	age := 20

	a, b := "Rajesh", 15
	a, b = "Ramesh", 20

	fmt.Println(foo, foo1)
	fmt.Println(bar)
	fmt.Println(i, j)
	fmt.Println(x, y)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(a, b)
}
