// Multiple Return Values

package main

import (
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}

	result := a / b

	return result, nil
}

func main() {

	fmt.Println(divide(10, 2))
}
