package main

import (
	"fmt"
	"reflect"
)

// type Object interface {
// }

type Person struct {
	name string
}

func main() {

	var a interface{}
	a = 5
	// var p interface{}
	// a = Person{"Rajesh"}

	fmt.Println(reflect.TypeOf(a).Name())

	if _, ok := a.(Person); ok {
		fmt.Println("It's a person")
	} else {
		fmt.Println("It's not a person")
	}

	// fmt.Println("is it a person?", a.(Person))

	switch a.(type) {
	case int:
		fmt.Println("It's int")
	case Person:
		fmt.Println("It's a person")
	default:
		fmt.Println("Not matching any!")
	}

}
