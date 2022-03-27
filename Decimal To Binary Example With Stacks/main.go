package main

import (
	"fmt"
	"log"
	"strconv"
)

const STACK_SIZE int = 10

type stack struct {
	data [10]int
	top  int
}

func main() {

	var arr [STACK_SIZE]int

	var n stack
	n.top = -1
	n.data = arr

	var InputNumber int
	var binaryNumber string

	fmt.Printf("Please enter a number: ")
	fmt.Scanf("%d\n", &InputNumber)

	for InputNumber >= 1 {
		remain := InputNumber % 2
		InputNumber = InputNumber / 2
		push(&n, remain)
	}

	for peek(&n) != -1 {
		binaryNumber = binaryNumber + strconv.Itoa(pop(&n))
	}

	fmt.Println(binaryNumber)

}

func push(stk *stack, numberToAdd int) {
	if stk.top == STACK_SIZE-1 {
		log.Fatal("STACK DOLU")
	} else {
		stk.top = stk.top + 1
		stk.data[stk.top] = numberToAdd
	}
}

func pop(stk *stack) int {
	if stk.top == -1 {
		fmt.Printf("Stack is empty\n")
	} else {
		stk.top = stk.top - 1
		return stk.data[stk.top+1]
	}

	return -1
}

func peek(stk *stack) int {
	if stk.top == -1 {
		return -1
	} else {
		return stk.data[stk.top]
	}
}
