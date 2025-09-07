package main

import (
	"fmt"
	"sync"
)

// Producer
func producer(id int, count int, buffer chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		item := id*100 + i // unique item ID
		buffer <- item
		fmt.Printf("Produced: %d by Producer-%d\n", item, id)
	}
}

// Consumer
func consumer(id int, buffer chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range buffer {
		fmt.Printf("Consumed: %d by Consumer-%d\n", item, id)
	}
}

// Entry point to run the producer-consumer system
func StartProductionConsumption(numProducers, numConsumers, itemsPerProducer, bufferSize int) {
	buffer := make(chan int, bufferSize)
	var wg sync.WaitGroup

	// Start all producers
	for i := 0; i < numProducers; i++ {
		wg.Add(1)
		go producer(i, itemsPerProducer, buffer, &wg)
	}

	// Start all consumers
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go consumer(i, buffer, &wg)
	}

	// Wait for all consumers to finish
	wg.Wait()
	close(buffer)

}

// Main function to test
func main() {
	StartProductionConsumption(2, 3, 5, 4)
}
