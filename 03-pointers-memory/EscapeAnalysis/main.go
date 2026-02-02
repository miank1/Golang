// Escape analysis is the concept whether a variable stored in the stack or move to heap area.
// Keep variables on the stack whenever possible to reduce GC pressure and improve performance.
// Escape analysis is a compile-time optimization where the Go compiler decides
// whether variables can be allocated on the stack or must escape to the heap based on their lifetime.

package main

// Escapes to heap
func getPointer() *int {
	x := 10
	return &x
}

// stays on stack
func sum() int {
	x := 10
	return x
}

// Slices and map usually escape
// Slices header on stack - backing array on heap

func makeSlice() []int {
	s := []int{1, 2, 3}
	return s
}

// escapes to heap due to returned pointer.
func bad() *int {
	x := 10
	return &x
}

// to check escape analysis - go build -gcflags="-m"
