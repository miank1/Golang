// x is a local variable. Normally, locals live on the stack and get cleaned up when the function exits.
// But here you return its address (&x).
// If Go kept x on the stack, its memory would vanish after the function returns → invalid pointer.
// Go’s compiler detects this and makes x escape to the heap instead, so it survives after function return.
// Go automatically decides stack vs heap allocation using escape analysis at compile time.
// If a variable’s lifetime is proven to end inside the function, it stays on the stack.
// If it “escapes” (like returning &x), it’s allocated on the heap.

package main

import "fmt"

func makeLocal() *int {
	x := 42
	return &x
}

func main() {
	result := makeLocal()

	fmt.Println("Result is ", *result)

}
