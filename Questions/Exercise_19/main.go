// Fan-In Pattern (Merging Channels)

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	s1 := []string{"A1", "A2", "A3"}
	s2 := []string{"B1", "B2", "B3"}
	out := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(200 * time.Millisecond)
			ch1 <- s1[i]
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(200 * time.Millisecond)
			ch2 <- s2[i]
		}
		close(ch2)
	}()

	go func() {
		for ch1 != nil || ch2 != nil {
			select {
			case msg, ok := <-ch1:
				if !ok {
					ch1 = nil // disable this case
					continue
				}
				out <- msg

			case msg, ok := <-ch2:
				if !ok {
					ch2 = nil // disable this case
					continue
				}
				out <- msg
			}
		}
		close(out)
	}()

	for x := range out {
		fmt.Println(x)
	}

}
