package main

import (
	"fmt"
	"sync"
)

func sendData(dataChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		dataChannel <- i
		fmt.Printf("Sent: %d\n", i)
	}
	close(dataChannel)
}

func receiveData(dataChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range dataChannel {
		fmt.Printf("Received : %d\n", num)
	}
}

func main() {
	dataChannel := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go sendData(dataChannel, &wg)
	go receiveData(dataChannel, &wg)

	wg.Wait()
}
