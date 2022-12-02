package main

import (
	"fmt"

	"github.com/rajesh4b8/go-learn-batch-5/examples/helloworld/print"
)

func main() {
	fmt.Println("Hello World!")
	print.PrintSomething()
	// print.getVersion() // can't access here because it's not exported - starting with lower case
}
