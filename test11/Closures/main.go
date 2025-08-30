package main

import "fmt"

func safeDivide(a, b int) int {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if b == 0 {
		panic("divide by zero")
	}

	return a / b

}

func main() {
	fmt.Println(safeDivide(10, 2)) // 5
	fmt.Println(safeDivide(10, 0)) // 0, with recovery message
}
