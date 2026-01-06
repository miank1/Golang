package main

import "fmt"

func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func main() {
	result, ok := divide(10, 2)
	if ok {
		fmt.Println(result)
	}

}
