// fanOut - Multiple Go routines producing data
// FanIn - Single Go routine consuming and aggregating data

package main

import (
	"fmt"
	"sync"
)

func SendNumber_1(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		ch <- i
	}

	//close(ch)
}

func SendNumber_2(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 4; i < 5; i++ {
		ch <- i
	}

	//close(ch)
}
func main() {

	ch := make(chan int, 5)

	var wg sync.WaitGroup

	wg.Add(2)
	go SendNumber_1(ch, &wg)
	go SendNumber_2(ch, &wg)

	for x := range ch {
		fmt.Println(x)
	}

	wg.Wait()

	close(ch)

}
