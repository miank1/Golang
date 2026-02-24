package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "World"

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	close(ch)

	ch1 := make(chan int, 3)
	ch1 <- 10
	ch1 <- 20
	ch1 <- 30

	close(ch1)

	for x := range ch1 {
		fmt.Println(x)
	}

	ch2 := make(chan string, 1)
	ch2 <- "hello"

	fmt.Println(<-ch2)
}
