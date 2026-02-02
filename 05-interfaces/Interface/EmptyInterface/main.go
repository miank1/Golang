// All type has zero methods so they satisfy empty interface.

package main

import "fmt"

func PrintAnything(v interface{}) {
	fmt.Println(v)
}

func main() {
	PrintAnything("Hello World !!")
	PrintAnything(2)
	PrintAnything(true)

	var data []interface{}
	data = append(data, 1, "hello", true)

}
