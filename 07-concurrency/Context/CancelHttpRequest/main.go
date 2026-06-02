package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	result, err := callPaymentService(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusGatewayTimeout)
		return
	}

	fmt.Println(w, result)
}

func callPaymentService(ctx context.Context) (string, error) {
	select {
	case <-time.After(5 * time.Second):
		return "payment success", nil
	case <-ctx.Done():
		return "Payment is done", ctx.Err()
	}
}

func main() {
	http.HandleFunc("/pay", paymentHandler)
	http.ListenAndServe(":8080", nil)
}
