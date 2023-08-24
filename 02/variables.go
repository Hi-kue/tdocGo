package main

import (
	"fmt"
)

func main() {
	// Go allows for the use of multiple variable declarations at once
	var a, b, c int = 10, 20, 50

	fmt.Printf("a:%d, b:%d, c:%d", a, b, c)

	// Go allows variables without a specified type, the type is inferred during compile time
	var d, e, f = true, 2.3, "four"

	fmt.Printf("d:%t, e:%f, f:%s", d, e, f)

	// The := is shorthand for declaring and initializing a variable
	g, h, i := 1, 2.3, "four"
	fmt.Printf("g:%d, h:%f, i:%s", g, h, i)
}