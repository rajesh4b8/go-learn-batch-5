package main

import (
	"fmt"
	"unsafe"
)

// empty structs -> no properties inside
// it takes 0 bytes
type emty struct {
}

type person struct {
	name string
	age  int
}

type ExchangeService struct {
}

type CurrencyConverterService struct {
	exService ExchangeService
}

func (c CurrencyConverterService) convert() {
	fmt.Println("Inside convert!!")

	// c.exService.DoSomething()
}

func main() {
	e := emty{}
	fmt.Println(e)

	// define and declare a struct
	// struct in-line definition
	p := struct {
		name string
		age  int
	}{}
	p.name = "Rajesh"
	p.age = 12

	q := struct {
		name string
		age  int
	}{}
	q.name = "Rajesh"
	q.age = 12

	fmt.Println(p)

	i := 125555
	name := "Rajesh"

	fmt.Printf("Size of i: %v and type: %T\n", unsafe.Sizeof(i), i)
	fmt.Println("Size of name", unsafe.Sizeof(name))
	fmt.Println("Size of empty struct", unsafe.Sizeof(e))
	fmt.Println("Size of empty struct", unsafe.Sizeof(struct{}{}))
	fmt.Println("Size of bool", unsafe.Sizeof(true))

	// in-line defenition for empty struct
	x := struct{}{}
	y := struct{}{}

	fmt.Printf("x at %p & y at %p\n", &x, &y)

	fmt.Println("x == y", x == y)
	fmt.Println("&x == &y", &x == &y)
	fmt.Println("p == q", p == q)
	fmt.Println("&p == &q", &p == &q)

	service := CurrencyConverterService{
		ExchangeService{},
	}
	service.convert()

}
