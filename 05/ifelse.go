package main

/*
	If/Else Statements in Go

	Notes:
		- If/else statements are conditional statements, like any other language
		- If/else statements can be used in conjunction with for loops, as well as switch statements
		- If/else statements can also be used to check for errors, which is a common use case
*/

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Enter a number between 1 & 10: ")
	fmt.Scan(&num)

	// Check if num in range between 1 & 10
	if num >= 1 && num <= 10 {
		fmt.Print("Number is in range 1-10\n")
	} else {
		fmt.Print("Number is not in range between 1-10\n")
	}

	// Using IsInRange 
	if IsInRange(num) {
		fmt.Print("Number is in range of 1-10\n")
	} else {
		fmt.Print("Number is not in range between 1-10\n")
	}
}

func IsInRange(num int) bool {
	return num >= 1 && num <= 10
}
