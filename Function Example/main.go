package main

// imports
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin) // Create reader

	input, err := reader.ReadString('\n') // Take strings
	if err != nil {                       // Check statement
		os.Exit(-1)
	}

	input = strings.TrimSpace(input) // Remove blanks

	inputInteger, err := strconv.Atoi(input) // Convert integer
	if err != nil {                          // Check statement
		os.Exit(-1)
	}

	result, err := controlAge(int64(inputInteger)) // Call function
	if err != nil {                                // Check statement
		fmt.Println(err)
		os.Exit(-1) // exit program with -1
	}

	fmt.Printf("Age Control: %t", result) // Print result
}

// function
// take int64
// return bool and error
func controlAge(age int64) (bool, error) {

	if age < 0 {
		err := errors.New("age can not be negative")
		return false, err // return false and error that created by me
	}

	return true, nil // return true and null
}
