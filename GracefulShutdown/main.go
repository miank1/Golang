package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// A background worker that listens to ctx.Done()
func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("🛑 Worker %d stopping gracefully...\n", id)
			return
		default:
			fmt.Printf("⚙️ Worker %d is processing...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Create cancelable context
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// Start multiple workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	// Listen for system interrupt (Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop // Block until a signal is received
	fmt.Println("\n🧠 Received shutdown signal!")
	cancel() // Tell all workers to stop

	wg.Wait()
	fmt.Println("✅ All workers have exited gracefully.")
}
