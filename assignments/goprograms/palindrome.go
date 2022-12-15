package main

import "fmt"

func main() {

	var palNum, remainder int

	fmt.Print("Enter the Number to check Palindrome = ")
	fmt.Scanln(&palNum)

	reverse := 0

	for temp := palNum; temp > 0; temp = temp / 10 {
		remainder = temp % 10
		reverse = reverse*10 + remainder
	}

	fmt.Println("The Reverse of the Given Number = ", reverse)
	if palNum == reverse {
		fmt.Println(palNum, " is a Palindrome Number")
	} else {
		fmt.Println(palNum, " is Not a Palindrome Number")
	}
}
