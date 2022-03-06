package main

import "fmt"

func main() {

	// Array Decleration type 1
	// numbers := [5] int {20,30,50,2,1}

	/*
		// Array Decleration type 2
		var numbers [5] int
		numbers[0] = 20
		numbers[1] = 30
		numbers[2] = 50
		numbers[3] = 2
		numbers[4] = 1
	*/

	numbers := [5]int{20, 30, 50, 2, 1}

	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				temp := numbers[j]
				numbers[j] = numbers[i]
				numbers[i] = temp
			}
		}
	}

	fmt.Println(numbers)

}
