package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("lets do")

	slice_1 := []int{1, 2, 3, 4, 5, 6}
	slice_2 := []int{11, 12, 13, 14, 15}
	fmt.Println(slice_1)

	fmt.Println(slice_2)

	// To append slice 2 to slice 1
	slice3 := append(slice_1, slice_2...)
	fmt.Println("Merging slice 1 and slice 2 ", slice3)

	// To delete element present at ith index
	i := 3
	slice_1 = append(slice_1[:i], slice_1[i+1:]...)
	fmt.Println("After deleting at position 3: ", slice_1)

	// Adding adding value 4 at to the slice_1 at position 3
	slice_1 = append(slice_1[:i], append([]int{4}, slice_1[i:]...)...)
	fmt.Println("Add 4 at to the slice_1 at position 3: ", slice_1)

	// Insert a new slice of length j at index i
	a := []int{1, 2, 3, 7, 8, 9, 10}
	b := []int{4, 5, 6}
	a = append(a[:3], append(b, a[i:]...)...)
	fmt.Println(a)

	//Pop element from stack
	fmt.Println(a[:len(a)-1])

	slice4 := []int{9, 5, 4, 3, 6, 2, 7}
	sort.Ints(slice4)
	fmt.Println(slice4)
	fmt.Println(sort.Reverse(sort.IntSlice(slice4)))

	// to send numbers to interface arguments
	add(1, 2, 3, 4, 5)

	// to send different data types to the function
	add("sda", "hi", 123, []int{1, 2, 3, 4}, [2]string{"santosh", "thota"})

}

func add(i ...interface{}) {
	fmt.Println(i...)

	for indx, val := range i {
		switch val.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", indx)
		case float64:
			fmt.Printf("param #%d is a float64\n", indx)
		case int, int64:
			fmt.Printf("param #%d is an int\n", indx)
		case nil:
			fmt.Printf("param #%d is nil\n", indx)
		case string:
			fmt.Printf("param #%d is a string\n", indx)

		default:
			fmt.Printf("param #%d's type is unknown\n", indx)
		}
	}

}
