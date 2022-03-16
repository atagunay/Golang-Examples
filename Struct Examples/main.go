package main

import "fmt"

type product struct {
	id   int
	name string
}

func main() {

	var temp product
	temp = product{1, "computer"}
	fmt.Println(temp)
	save(temp)
	temp.save2()

	temp2 := &temp
	fmt.Println(*temp2)
	temp2.save2()
	save(*temp2)

}

func save(p product) {
	fmt.Println("added with save")
}

func (p product) save2() {
	fmt.Println("added with save2")
}
