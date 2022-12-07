package main

import "fmt"

// struct is the keyword
// person is the name of the struct
// struct definition
type person struct {
	firstName string
	lastName  string
	age       int
	addresses []address
}

type address struct {
	city string
	zip  int
}

func main() {
	// declaration without initialization -> zero values for all the props
	// var p person

	// declare and init
	// var p person = person{"Rajesh", "Reddy", 12}
	// p := person{"Rajesh", "Reddy", 12}
	p := person{age: 12, firstName: "Rajesh"}
	p.age = 15

	addr1 := address{"Dallas", 25356}
	addr2 := address{zip: 64334, city: "Texas"}
	p.addresses = []address{addr1, addr2}

	p.printDetails()
}

func (p person) printDetails() {
	fmt.Printf("%+v\n", p)

	fmt.Println(p.addresses[0].city)
	for _, add := range p.addresses {
		fmt.Println(add.city)
	}
}
