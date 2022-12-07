package main

import "fmt"

func main() {
	a := 20
	fmt.Println("Before", a)

	doIncrement(a)

	fmt.Println("After", a)

	s := make([]int, 5, 10)
	s[0] = 1
	s[1] = 2
	addMoreToSlice(s)

	fmt.Println("slice1", "Length:", len(s), "Cap:", cap(s), "Contents:", s)

	// fmt.Println(s[5])
	arr := [5]int{1, 2, 3}
	addMoreToArray(arr)
	fmt.Println("Arr", arr)
}

func doIncrement(i int) {
	i = i + 1
	fmt.Println("doIncrement", i)
}

func addMoreToSlice(values []int) {
	// append
	values = append(values, 4)
	// values[3] = 4
	fmt.Println("slice1", "Length:", len(values), "Cap:", cap(values), "Contents:", values)
	fmt.Println("inside addMore", values)
	fmt.Println("inside addMore", values[5])
}

func addMoreToArray(values [5]int) {
	values[3] = 4
	fmt.Println("Inside addMoreToArray", values)
}
