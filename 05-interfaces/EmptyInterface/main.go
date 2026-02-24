package main

import "fmt"

// Any empty interface has no methods, by default all types implement the empty interface.

func main() {
	test("thisisstring")
	test(10)
	test(true)
}

func test(a interface{}) {
	fmt.Printf("%v, %T \n", a, a)
}
