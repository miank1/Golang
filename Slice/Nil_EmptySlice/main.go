// Nil and Empty slice

package main

import "fmt"

func main() {
	var a []int  // nil slice
	b := []int{} // empty slice

	fmt.Println(a == nil)
	fmt.Println(b == nil)
}
