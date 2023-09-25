package main

import "fmt"

// Student /* Singleton object
type Student struct {
	name string
}

// Variable for keeping single object's reference
var student *Student

// Constructor method
// Return address of student
func getNewStudent() *Student {
	if student == nil { // if student is not refer to any object create a student object
		student = &Student{name: "ata"}
	} else {
		fmt.Println("Object already existed")
	}
	return student
}

func main() {
	fmt.Println(&getNewStudent().name)
	fmt.Println(&getNewStudent().name)
	fmt.Println(&getNewStudent().name)
	fmt.Println(&getNewStudent().name)
}
