package main

import (
	"fmt"
	"math"
)

func main() {

	var meal_cost float64
	var tip_percent int32
	var tax_percent int32

	fmt.Printf("Please enter meal cost: ")
	fmt.Scanf("%f\n", &meal_cost)

	fmt.Printf("Please enter tip percent: ")
	fmt.Scanf("%d\n", &tip_percent)

	fmt.Printf("Please enter tax percent: ")
	fmt.Scanf("%d\n", &tax_percent)

	tip := meal_cost * float64(tip_percent) / 100
	tax := meal_cost * float64(tax_percent) / 100

	result := meal_cost + tip + tax
	fmt.Println(math.Round(result))

}
