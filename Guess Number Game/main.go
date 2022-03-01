package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var TRIAL_COUNT int = 10

func main() {

	generatedNumber := rand.Intn(100)
	fmt.Println(generatedNumber)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < TRIAL_COUNT; i++ {

		fmt.Println("Plese enter your guess...")
		userGuessString, _ := reader.ReadString('\n')

		userGuessString = strings.TrimSpace(userGuessString)
		userGuessInt, _ := strconv.Atoi(userGuessString)

		if userGuessInt == generatedNumber {
			fmt.Println("Guess is TRUE !")
			break

		} else {

			if userGuessInt < generatedNumber {
				fmt.Println("Your guess is lower than correct number")
			} else {
				fmt.Println("Your guess is bigger than correct number")
			}

		}

	}
}
