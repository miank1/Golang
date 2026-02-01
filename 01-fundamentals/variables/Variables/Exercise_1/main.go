// Write a small Go program that:

// Declares a global variable counter using var.

// In main(), declare a local variable counter using :=.

// Print both to show they are different.

// Add a function increment() that increases the global counter by 1.

// Call it twice and print the final global counter.

package main

import "fmt"

var counter int

func increment() {
	counter++
}

func main() {

	counter := 0

	fmt.Println("Local Variable: ", counter)
	increment()
	increment()

	fmt.Println("Global counter after increment:", counter) // still local
	fmt.Println("Global counter value:", getGlobalCounter())
}

func getGlobalCounter() int {
	return counter // returns global
}
