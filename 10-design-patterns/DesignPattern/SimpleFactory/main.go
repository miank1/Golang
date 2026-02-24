package main

// Centralized place for object creation

import "fmt"

type Door interface {
	GetDescription() string
}

type WoodenDoor struct{}

func (w *WoodenDoor) GetDescription() string {
	return "I'm a wooden door"
}

type IronDoor struct{}

func (i *IronDoor) GetDescription() string {
	return "I'm an iron door"
}

func DoorFactory(doorType string) Door {
	switch doorType {
	case "wood":
		return &WoodenDoor{}
	case "iron":
		return &IronDoor{}
	default:
		return nil
	}
}

func main() {
	door := DoorFactory("wood")
	if door != nil {
		fmt.Println(door.GetDescription())
	} else {
		fmt.Println("Unknown door type")
	}
}
