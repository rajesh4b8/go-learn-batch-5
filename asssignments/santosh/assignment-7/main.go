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
}

// advertisement structure
type Advertisement struct {
	addName   string
	noOfClick int
	clickRate int
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
}

// function over loading is not allowed in go
func (p TimeAndMaterial) calculateIncome() int {
	return p.hourlyRate * p.noOfHours
}

func (p TimeAndMaterial) source() string {
	return p.projectName
}

// implementing interface methods for advertisement
func (a Advertisement) calculateIncome() int {
	return a.noOfClick * a.clickRate
}
func (a Advertisement) source() string {
	return a.addName
}

func calculateNetIncome(steams []Income) int {
	var netIncome int
	for _, s := range steams {
		fmt.Printf("Income from %v : $%d\n", s.source(), s.calculateIncome())
		netIncome += s.calculateIncome()
	}

	return netIncome
}

func main() {
	p1 := FixedBilling{"prj 1", 20000}
	p2 := TimeAndMaterial{"prj 2", 100, 40}
	p3 := TimeAndMaterial{"prj 3", 2000, 25}

	a1 := Advertisement{"Dell Alienware advertisement", 10000, 2}

	fmt.Println(p1, p2, a1)
	fmt.Println("Total income", calculateNetIncome([]Income{p1, p2, p3, a1}))
}
