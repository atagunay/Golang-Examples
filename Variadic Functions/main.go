package main

import (
	"fmt"
)

func main() {

	count := 0
	var input int

	fmt.Println("Please enter your number count: ")
	fmt.Scanf("%d", &count)

	numbers := make([]int, 0)

	for i := 0; i < count; i++ {
		fmt.Printf("Enter a number: ")
		fmt.Scan(&input)
		numbers = append(numbers, input)
	}

	fmt.Println(add(numbers...))

}

func add(numbers ...int) int {
	result := 0

	for i := 0; i < len(numbers); i++ {
		result = result + numbers[i]
	}

	return result
}
