// wherther the code will compile or not ??

package main

import "fmt"

func main() {
	ch := make(chan int) //unbuffered -- creates a deadlock
	ch <- 10
	fmt.Println(<-ch)
}

// Fix would be add a go routine or use unbuffered channels.

// go func() {
// 	ch <- 10
// }()
// fmt.Println(<-ch)
