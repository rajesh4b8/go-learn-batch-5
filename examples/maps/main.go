package main

import "fmt"

func main() {
	// Delcare
	colors := make(map[string]string)
	colors["Red"] = "#536226"
	colors["White"] = "#111111"
	updateColor(colors)

	fmt.Println(colors)
	fmt.Println("Color code for Red is", colors["Red"])

	shapes := map[string]int{
		"Triangle":  3,
		"Rectangle": 4,
		"Square":    4,
	}

	// It overrides the existing value
	shapes["Square"] = 0

	// delete from map
	delete(shapes, "Square")

	fmt.Println(shapes)
}

// Map is a reference data type -> the underlying data is updated
func updateColor(c map[string]string) {
	c["Black"] = "#000000"
}
