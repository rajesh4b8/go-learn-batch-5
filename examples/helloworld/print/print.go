package print

import "fmt"

var (
	PRINT_VERSION = 20
	i             = 20
)

func PrintSomething() {
	fmt.Println("Print something with version", getVersion())
}

func getVersion() int {
	return PRINT_VERSION
}
