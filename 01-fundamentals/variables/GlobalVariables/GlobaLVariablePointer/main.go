package main

import "fmt"

var username *string

func main() {
	newName := "John Doe"
	new_Name := fmt.Sprintf("%s", newName)
	username = &new_Name
	fmt.Println(*username)
}
