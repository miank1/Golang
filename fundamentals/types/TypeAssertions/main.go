package main

import (
	"fmt"
	"reflect"
)

func main() {
	var value interface{} = "Hello World"

	var str string = value.(string)
	fmt.Println(str)

	fmt.Println(reflect.TypeOf(str))

	// this will panic as interface
	// does not have int type
	var value2 int = value.(int)

	fmt.Println(value2)

}
