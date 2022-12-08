package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{"Rajesh", 12}

	fmt.Printf("Person data :: %+v\n", p1)

	// `&` is used to get the pointer
	personPointer := &p1
	fmt.Printf("person pointer addr :: %p\n", personPointer)

	// p2 := person{}
	// p2p := &p2
	// new will allocate a memory to the zero values and return the pointer for it
	p2p := new(person)
	fmt.Println("p2p from new()", p2p)

	// `*` is used to get the value at the pointer
	valueAtPointer := *personPointer
	fmt.Printf("Person value at pointer :: %+v\n", valueAtPointer)

	p1.updateNameByValue()
	fmt.Printf("Person updated by value :: %+v\n", p1)

	// (&p1).updateNameByRef()
	p1.updateNameByRef() // pointer indirection
	fmt.Printf("Person updated by Ref :: %+v\n", p1)
}

// struct is a value type
func (p person) updateNameByValue() {
	p.name = "Suresh"
}

func (pp *person) updateNameByRef() {
	// (*pp).name = "Mahesh"
	pp.name = "Mahesh" // pointer indirection
	fmt.Printf("person pointer addr :: %p\n", pp)
	fmt.Println("person pointer", pp)
}
