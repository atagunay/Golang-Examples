package main

import "fmt"

func main() {
	// Add odd numbers in 1-10 with for range
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	total := 0

	for _, number := range numbers {
		if number%2 != 0 {
			total = total + number
		}
	}

	fmt.Println(total)
}
