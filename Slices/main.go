package main

import "fmt"

func main() {

	// Slices decleration
	numbers := make([]int, 0)
	numbers = append(numbers, 10)
	numbers = append(numbers, 20)
	numbers = append(numbers, 20)
	numbers = append(numbers, 30)
	numbers = append(numbers, 40)
	numbers = append(numbers, 50)

	// From arr[1] to arr[4]
	fmt.Println(numbers[1:5])

	numbersCopy := make([]int, len(numbers)) // Create a new slice
	copy(numbersCopy, numbers)               // Copy numbers to numbersCopy

	fmt.Println(numbersCopy)

}
