package main

import "fmt"

func main() {

	i := 32
	fmt.Printf("Type %T, value %v \n", i, i)

	f := float32(i)
	fmt.Printf("Type %T, value %v \n", f, f)

}
