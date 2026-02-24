package main

import (
	"fmt"
	"math/rand/v2"
)

type Transport interface {
	BookRide(destination string) string
	Cost() float64
	Type() string
}

// Step 2: Implement Car
type Car struct {
	DriverName string
}

func (c Car) BookRide(destination string) string {
	return fmt.Sprintf("Car booked to %s. Driver: %s", destination, c.DriverName)
}

func (c Car) Cost() float64 {
	return 150 + rand.Float64()*50
}

func (c Car) Type() string {
	return "Car"
}

// Step 3: Implement Bike
type Bike struct {
	Rider string
}

func (b Bike) BookRide(destination string) string {
	return fmt.Sprintf("Bike ride to %s. Rider: %s", destination, b.Rider)
}

func (b Bike) Cost() float64 {
	return 50 + rand.Float64()*30 // Between 50 and 80
}

func (b Bike) Type() string {
	return "Bike"
}

// Step 4: Implement Auto
type Auto struct {
	PlateNumber string
}

func (a Auto) BookRide(destination string) string {
	return fmt.Sprintf("Auto ride to %s. Plate: %s", destination, a.PlateNumber)
}

func (a Auto) Cost() float64 {
	return 80 + rand.Float64()*40 // Between 80 and 120
}

func (a Auto) Type() string {
	return "Auto"
}

func StartBooking(t Transport, destination string) {
	fmt.Println("🚗 Booking Summary")
	fmt.Println("-------------------")
	fmt.Println("Ride Type  :", t.Type())
	fmt.Println("Destination:", destination)
	fmt.Printf("Fare       : ₹%.2f\n", t.Cost())
	fmt.Println()
}

func main() {

	car := Car{
		DriverName: "Rahul",
	}
	bike := Bike{
		Rider: "Amit",
	}
	auto := Auto{
		PlateNumber: "AA12 ZY1234",
	}

	tranportOptions := []Transport{car, bike, auto}
	for _, t := range tranportOptions {
		StartBooking(t, "MG Road")
	}
}
