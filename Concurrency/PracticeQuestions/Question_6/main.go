// Fan In - Concurrency Pattern question

package main

import "fmt"

func Producer_1(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func Producer_2(ch chan<- string) {
	for i := 5; i < 10; i++ {
		ch <- fmt.Sprintf("String %d", i)
	}
	close(ch)
}

func FanIn(ch1 <-chan int, ch2 <-chan string) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for ch1 != nil || ch2 != nil {
			select {
			case v, ok := <-ch1:
				if !ok {
					ch1 = nil
					continue
				}
				out <- v
			case v, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				}
				out <- v
			}
		}
	}()
	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go Producer_1(ch1)
	go Producer_2(ch2)

	fanIn := FanIn(ch1, ch2)

	for v := range fanIn {
		fmt.Println(v)
	}
}
