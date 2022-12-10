package main

import "fmt"

// interace is the keyword to define any interface
type Income interface {
	calculateIncome() int
	source() string
}

// to implement `Income` you must simply implement all the methods
type FixedBilling struct {
	projectName string
	bidAmount   int

	// methods
	// calculateIncome()
	// source()
}

// implementing the first method `calculateIncome` of `Income` interface
func (p FixedBilling) calculateIncome() int {
	return p.bidAmount
}

// implementing the second method `source` of `Income` interface
func (p FixedBilling) source() string {
	return p.projectName
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int

	// methods
	// calculateIncome() string
	// source() int
	// noOfHours() int
}

func (p TimeAndMaterial) calculateIncome() int {
	return p.hourlyRate * p.noOfHours
}

func (p TimeAndMaterial) source() string {
	return p.projectName
}

func (p TimeAndMaterial) getNoOfHours() int {
	return p.noOfHours
}

// Advertisement income

func main() {
	p1 := FixedBilling{"prj 1", 20000}
	p2 := TimeAndMaterial{"prj 2", 100, 40}
	p3 := TimeAndMaterial{"prj 3", 2000, 25}
	fmt.Println("no of hours for p3", p3.getNoOfHours())

	fmt.Println(p1, p2)
	fmt.Println("Total income", calculateNetIncome([]Income{p1, p2, p3}))
}

func calculateNetIncome(steams []Income) int {
	var netIncome int
	for _, s := range steams {
		fmt.Printf("Income from %v : $%d\n", s.source(), s.calculateIncome())
		netIncome += s.calculateIncome()
	}

	return netIncome
}
