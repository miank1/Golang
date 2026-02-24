// Slice shares a shared memory

package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4}
	a := arr[:2]
	b := arr[2:]

	a[1] = 99
	fmt.Println(arr)
	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("arr - ", arr)

}
