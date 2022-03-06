package main

import "fmt"

func main() {

	// KEY - VALUE
	dictionary := make(map[string]string) // decleration type 1
	//dictionary := map[string]string{"glass": "bardak","car": "araba"} // decleration type 2

	dictionary["book"] = "Kitap"
	dictionary["table"] = "Masa"
	dictionary["phone"] = "Telefon"
	dictionary["computer"] = "Bilgisayar"

	fmt.Println("Map: ", dictionary)
	fmt.Println("Element numbers: ", len(dictionary))

	delete(dictionary, "phone")
	fmt.Println("Element numbers: ", len(dictionary))

	value, isHere := dictionary["table"]
	fmt.Println("Value of table: ", value)
	fmt.Println("Is here: ", isHere)

}
