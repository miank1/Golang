package main

import (
	"fmt"
	paymentsystem "interfaceExample/PaymentSystem"
)

// Both Dog and Cat satisfy the Speaker interface because they both have a Speak() method.

// Define an interface
type Speaker interface {
	Speak()
}

// Define a struct
type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

func main() {
	var s Speaker
	s = Dog{}
	s.Speak() // Output : Woof!

	s = Cat{}
	s.Speak() // Output : Meow!

	stripe := paymentsystem.Stripe{}
	paymentsystem.MakePayment(stripe, 100.0)
	paypal := paymentsystem.Stripe{}
	paymentsystem.MakePayment(paypal, 50.0)
}
