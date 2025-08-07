package main

import (
	"fmt"
	"time"
)

func generate(msg string) <-chan string {
	ch := make(chan string)

	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return ch
}

func fanIn(input1, input2 <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for {
			out <- <-input1
		}
	}()

	go func() {
		for {
			out <- <-input2
		}
	}()

	return out
}

func main() {

	ch1 := generate("Apple")
	ch2 := generate("Banana")

	combined := fanIn(ch1, ch2)

	for i := 0; i < 10; i++ {
		fmt.Println(<-combined)
	}

}
