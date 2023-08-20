package main

// default imports
import (
	"fmt"
)

func main() {
	// var name *string -> nil by default, pointer to string
	var name *string

	fmt.Println("Type in your name:")
	fmt.Scanln(&name)

	// checking if name is nil
	if name != nil {
		fmt.Println("Hello", name, " and Hello World!")
	} else {
		fmt.Println("Hello World!")
	}
}
