package main

import "fmt"

// Vechile is an interface that all the vechile types must implement

type Vechile interface {
	Drive() string
	GetInfo() string
}

type Car struct {
	Brand string
}

func (c Car) Drive() string {
	return fmt.Sprintf("🚗 %s car is driving!", c.Brand)
}

func (c Car) GetInfo() string {
	return fmt.Sprintf("Car Brand: %s", c.Brand)
}

type Bike struct {
	Model string
}

func (b Bike) Drive() string {
	return fmt.Sprintf("🏍️ %s bike is riding!", b.Model)
}

func (b Bike) GetInfo() string {
	return fmt.Sprintf("Car Brand: %s", b.Model)
}

func NewVehicle(vechileType, details string) Vechile {
	if vechileType == "car" {
		return Car{Brand: details}
	} else if vechileType == "bike" {
		return Bike{Model: details}
	}
	return nil
}

func main() {
	car := NewVehicle("car", "Tesla")
	if car != nil {
		fmt.Println(car.Drive())   // 🚗 Tesla car is driving!
		fmt.Println(car.GetInfo()) // Car Brand: Tesla
	} else {
		fmt.Println("Invalid vehicle type")
	}

	// Create a Bike with a model
	bike := NewVehicle("bike", "Yamaha")
	if bike != nil {
		fmt.Println(bike.Drive())   // 🏍️ Yamaha bike is riding!
		fmt.Println(bike.GetInfo()) // Bike Model: Yamaha
	} else {
		fmt.Println("Invalid vehicle type")
	}

	// Handling an unknown vehicle type
	unknown := NewVehicle("truck", "Ashok")
	if unknown == nil {
		fmt.Println("Unknown vehicle type!")
	}
}
