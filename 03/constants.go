package main

import (
	"fmt"
	"math"
)

func main() {
	// Go constants are declared using the `const` keyword
	// Similar to variables, constants can be of any type or no type and inferred during compile time
	const a = 10
	fmt.Printf("a:%d", a)

	// Constant expressions perform arithmetic w/arbitrary precision
	const n = 500000000
	const d = 3e20 / n

	fmt.Printf("d:%f", d)

	// Numeric constants have no type until it's given one, by explicit conversion
	fmt.Printf("int64(d):%d", int64(d))

	fmt.Printf("math.Sin(n):%f", math.Sin(n))

	arithmetic()
}

func arithmetic() {
	// Calculating the hypotenuse of a triangle
	const base = 2

	const (
		a = 10 // Side `a`
		b = 20 // Side `b`
	)

	var c float64 = math.Sqrt(math.Pow(a, base) + math.Pow(b, base))
	fmt.Printf("c:%f", c)
}
