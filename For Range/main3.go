package main

import "fmt"

func main() {

	// visit map
	dictionary := make(map[string]string)
	dictionary["book"] = "Kitap"
	dictionary["table"] = "Masa"
	dictionary["phone"] = "Telefon"
	dictionary["computer"] = "Bilgisayar"

	for key, value := range dictionary {
		fmt.Printf("%s value is: %s \n", key, value)
	}
}
