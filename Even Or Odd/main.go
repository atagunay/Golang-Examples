package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Please Enter a number") // Inform user

	// Create a reader.
	// This reader will listen keyboard. Because keyboard is standart input for pc (os.Stdin)
	reader := bufio.NewReader(os.Stdin)

	// take input and return string when user press enter
	// In GO, functions can return different types
	// this function returns string and error
	// it is similar to exception
	input, err := reader.ReadString('\n')

	if err != nil { // Control block, if erorr occurs, it will be print out
		log.Fatal(err) // print error
	}

	fmt.Println("input", input)                           // print input
	fmt.Println("type of input: ", reflect.TypeOf(input)) // print type of input

	// trim the spaces
	// you can also use this
	input = strings.TrimSpace(input)

	numericResult, err := strconv.Atoi(input) // Convert string to integer
	if err != nil {                           // Control block, if erorr occurs, it will be print out
		log.Fatal(err) // print error
	}

	fmt.Println("numericResult", numericResult)                         // print numericResult
	fmt.Println("type of numericResult", reflect.TypeOf(numericResult)) // print type of numericResult

	if numericResult%2 == 0 {
		fmt.Println("this is even")
	} else {
		fmt.Println("this is odd")
	}

}
