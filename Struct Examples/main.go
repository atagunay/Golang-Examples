package main

import "fmt"

func main() {
	type product struct {
		id   int
		name string
	}

	var temp product
	temp = product{1, "computer"}
	fmt.Println(temp)

	temp2 := &temp
	fmt.Println(*temp2)

}
