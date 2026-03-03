package main

import "fmt"

type PaymentProcessor interface {
	Process(amount float64) error
}

type CreditCard struct{}

type UPI struct{}

func (p *CreditCard) Process(amount float64) error {
	fmt.Printf("Processing %.2f via Credit Card\n", amount)
	return nil
}

func (p *UPI) Process(amount float64) error {
	fmt.Printf("Processing %.2f via UPI\n", amount)
	return nil
}

func main() {
	var p PaymentProcessor

	p = &CreditCard{}
	p.Process(1000)

	p = &UPI{}
	p.Process(500)
}
