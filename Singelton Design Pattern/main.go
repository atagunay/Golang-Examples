package main

import (
	"fmt"
	"sync"
)

// Student /* Singleton object
type Student struct {
	name string
}

// Variable for keeping single object's reference
var student *Student
var lock = &sync.Mutex{}

// Constructor method
// Return address of student
func getNewStudent() *Student {
	if student == nil { // if student is not refer to any object create a student object
		lock.Lock()
		defer lock.Unlock()
		// If two threads arrive lock.Lock() method, only one of them can enter inside
		// Other one waits for first one. When first thread complete its job, second one enter inside
		// This if condition blocks second thread to create same object again
		if student == nil {
			student = &Student{name: "ata"}
		} else {
			fmt.Println("Object already existed")
		}
	}
	return student
}

func main() {
	fmt.Println(&getNewStudent().name)
	fmt.Println(&getNewStudent().name)
	fmt.Println(&getNewStudent().name)
	fmt.Println(&getNewStudent().name)
}
