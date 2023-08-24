package main

import (
	"fmt"
)

func main() {
	// values are simple to understand in go, they are just like in any other language
	fmt.Println("Go" + "lang")                // Golang
	fmt.Printf("1 + 10 = %d", 1+10)           // 11
	fmt.Printf("%d / %d = %d", 10, 10, 10/10) // 1

	// this is similar to discrete math, T ^ F = T
	fmt.Println(true && false) // false (and)

	// this is similar to discrete math, T v F = T
	fmt.Println(true || false)                           // true (or)
	fmt.Println(!true)                                   // false (not)
	fmt.Printf("%s -> %s, %s -> %s", "T", "F", "F", "T") // T -> F, F -> T
}