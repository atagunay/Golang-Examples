package main

import "fmt"

func main() {

	number1 := 0
	fmt.Printf("Plase enter first number: ")
	fmt.Scanln(&number1)

	number2 := 0
	fmt.Printf("Plase second first number: ")
	fmt.Scanln(&number2)

	resultForNumber1 := 0
	for i := 1; i < number1; i++ {
		if number1%i == 0 {
			resultForNumber1 = resultForNumber1 + i
		}
	}

	resultForNumber2 := 0
	for i := 1; i < number2; i++ {
		if number2%i == 0 {
			resultForNumber2 = resultForNumber2 + i
		}
	}

	if resultForNumber1 == number2 && number1 == resultForNumber2 {
		fmt.Println("These are friend numbers")
	} else {
		fmt.Println("These are not friend numbers")
	}
}
