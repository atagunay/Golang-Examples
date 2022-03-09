package main

import "fmt"

func main() {

	sehirler := [3]string{"Ankara", "İstanbul", "İzmir"}

	for i, sehir := range sehirler {

		fmt.Printf("%d index is: %s \n", i, sehir)
	}

}
