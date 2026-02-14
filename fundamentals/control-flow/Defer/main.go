package main

import "fmt"

// defer executes when sorrounding function completes.

func main() {
	defer fmt.Println("This is printed once the function completes")

	fmt.Println("This is printed first")
	fmt.Println("This is printed second")

	// defer is scoped to the function which it is declared.
	greet()
	fmt.Println("This is printed after the greet is called.")

	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")

	fmt.Println("starting my Function...")
}

func greet() {
	defer fmt.Println("Printed after the first greeting")
	fmt.Println("first greeting")
}
