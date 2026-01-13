// Select With Default Case

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 10
	}()

	for i := 1; i <= 3; i++ {

		select {
		case x := <-ch:
			fmt.Println("Received ", x)
		default:
			fmt.Println("No data yet")
			time.Sleep(500 * time.Millisecond)
		}
	}

}
