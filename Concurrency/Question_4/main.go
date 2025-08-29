// You have two channels:

// ch1: sends numbers from 1 to 5.

// ch2: sends letters from A to E.

// A goroutine should write numbers into ch1.

// Another goroutine should write letters into ch2.

// The main goroutine should use a select {} loop to read from both channels and print values until both are closed.

package main

import (
	"fmt"
	"sync"
)

func Writer_1(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func Writer_2(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	str := []string{"A", "B", "C", "D", "E"}

	for i := 0; i < len(str); i++ {
		ch <- str[i]
	}
	close(ch)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go Writer_1(ch1, &wg)
	go Writer_2(ch2, &wg)

	done1, done2 := false, false
	for !(done1 && done2) {
		select {
		case val, ok := <-ch1:
			if ok {
				fmt.Println("Received from ch1:", val)
			} else {
				done1 = true
			}
		case val, ok := <-ch2:
			if ok {
				fmt.Println("Received from ch2:", val)
			} else {
				done2 = true
			}
		default:
			fmt.Println("No more data to read.")
		}
	}
	wg.Wait()
}
