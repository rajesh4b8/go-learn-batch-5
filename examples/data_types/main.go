package main

import "fmt"

func main() {
	var flag bool = true
	var str string = "Some string"
	var i int = -1234       // int, int8, int16, int32, int64
	var ui uint = 123       // uint, uint8, uint16, uint32, uint64, uintptr
	var r rune = 65346      // rune =~ int32
	var b byte = 255        // byte =~ uint8 0 to 255
	var c int8 = -127       // -127 to 127
	var f float32 = 43.6464 // float32, float64

	fmt.Println(flag)
	fmt.Println(str)
	fmt.Println(i)
	fmt.Println(ui)
	fmt.Println(r)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(f)
}
