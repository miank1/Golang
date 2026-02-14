package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	err := processRequest(ctx)

	if err != nil {
		http.Error(w, "Request cancelled: "+err.Error(), http.StatusRequestTimeout)
		return
	}

	fmt.Fprintln(w, "Request completed successfully!")
}

func processRequest(ctx context.Context) error {
	select {
	case <-time.After(4 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {

	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
