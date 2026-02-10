// Modifying elements
// Slice header is copied and underlying array is shared.
// Element modification is reflected everywhere.

package main

import "fmt"

func modify(s []int) {
	s[0] = 99
}

func main() {
	arr := []int{1, 2, 3}
	modify(arr)
	fmt.Println(arr)
}
