package main

// default imports
import (
	"fmt"
)

func main() {
	// var name, hobby : string
	var name, hobby string

	fmt.Println("Type in a name and a hobby: ")
	fmt.Scanln(&name, &hobby)

	// check if values provided is empty
	if name == "" || hobby == "" {
		fmt.Println("You did not provide a name or a hobby.")
	}

	var formatInput = fmt.Sprintf("Hello, my name is %s and I like %s.", name, hobby)
	fmt.Println(formatInput)
}
