package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{
		1, 2, 3, 4, 5, 6,
	}

	i := 1
	for i <= 4 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 10; j++ {
		random := rand.Intn(j)
		arr = append(arr, random)
	}

	fmt.Println(arr)


	// for loop to only get odd numbers from the array
	for _, value := range arr {
		if value % 2 == 0 {
			continue
		}
		fmt.Println(value)
	}

	// this is another variation to the above for loop
	for i := 0; i < len(arr); i++ {
		// check whether the number is even
		if arr[i] % 2 == 0 {
			continue
		}
		// print only odd numbers
		fmt.Println(arr[i])
	}
}
