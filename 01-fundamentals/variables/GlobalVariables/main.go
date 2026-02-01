package main

import (
	"fmt"
	"hello/GlobalVariables/util"
)

// Global Variable

var username string
var counter int

func main() {
	username = "John Doe"
	fmt.Println("Username is", username)

	// Access variable from across package
	fmt.Println("The path is: " + util.Mypath)

	for i := 0; i < 1000; i++ {
		counter += 1
	}
	fmt.Println("Counter value is ", counter)
}
