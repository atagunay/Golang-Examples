package main

import (
	"fmt"
	"os"
)

const START_INDEX = 2

func main() {

	fmt.Printf("Plase enter a number: ")

	number := 0
	fmt.Scanf("%d", &number)

	for i := START_INDEX; i < number; i++ {
		if number%i == 0 {
			fmt.Println("this is not a prime number")
			os.Exit(1)
		}
	}

	fmt.Println("this is a prime number")

}
