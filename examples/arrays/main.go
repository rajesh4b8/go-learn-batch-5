package main

import "fmt"

func main() {
	// Create array
	// the size of the array is 10
	// initialized with zero values of that type
	// var ai [10]int
	ai := [10]int{}
	fmt.Println(ai)

	ai[0] = 20
	ai[1] = 30
	ai[2] = 50

	fmt.Println(ai)

	as := [10]string{"Rajesh", "Ramesh", "Krishna"}
	as[4] = "Teja"
	fmt.Println(as)

	arrayWithValues := [5]int{0: 3, 4: 56}
	fmt.Println(arrayWithValues)

	// Slice of array from index 1 to index 4
	slice1 := ai[1:4]
	fmt.Println("slice1", "Length:", len(slice1), "Cap:", cap(ai), "Contents:", slice1)
}
