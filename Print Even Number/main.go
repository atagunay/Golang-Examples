package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// User inform about enter the number
	fmt.Println("Please enter a number")

	reader := bufio.NewReader(os.Stdin) // create reader
	input, _ := reader.ReadString('\n') // take strings

	input = strings.TrimSpace(input)       // remove space begin and end of the string
	inputNumeric, _ := strconv.Atoi(input) // strings convert to int

	for i := 0; i < inputNumeric; i++ { // for loop
		if i%2 == 0 { // control even or odd
			fmt.Println(i) // if i is even, print

		}
	}

}
