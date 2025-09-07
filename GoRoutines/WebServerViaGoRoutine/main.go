package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request:", r.URL.Path)
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello, World !!"
	}()

	// Receive the result from the channel
	result := <-ch
	w.Write([]byte(result))
	close(ch)
}
func main() {
	http.HandleFunc("/hello", hello)

	// Start the server at 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}
