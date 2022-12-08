package main

import "fmt"

type person struct {
	name string
}

func main() {
	var j int
	p := &j

	q := new(int)

	fmt.Println("p", p, "q", q)
	fmt.Println("*p", *p, "*q", *q)

	map1 := make(map[string]int)
	map1["triangle"] = 3

	fmt.Println(map1)

	// new is going init map with zero value (zero value for map is nil)
	// and return the pointer to it
	map2 := new(map[string]int)
	fmt.Println(map2)
	// (*map2)["traingle"] = 3 // the value at the pointer is nil -> because the zero value is nil
	fmt.Println(map2)

	// init with zero value and return pointer
	p1 := new(person)
	p1.name = "Rajesh" // pointer indirection
	fmt.Println(*p1)   // reading from pointer
}
