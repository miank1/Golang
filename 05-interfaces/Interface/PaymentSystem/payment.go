package paymentsystem

import "fmt"

// Real-world analogy:
// Imagine PaymentProcessor is like a plug type, and MakePayment is a wall socket that accepts any plug that fits.
// As long as the plug (i.e., the object p) has the correct shape (i.e., implements ProcessPayment), it can be used.

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

type Stripe struct{}

func (s Stripe) ProcessPayment(amount float64) error {
	fmt.Printf("Processing $%.2f via Stripe \n", amount)
	return nil
}

type PayPal struct{}

func (p PayPal) ProcessPayment(amount float64) error {
	fmt.Printf("Processing $%.2f via PayPal \n", amount)
	return nil
}

// Important
func MakePayment(p PaymentProcessor, amount float64) {
	p.ProcessPayment(amount)
}
