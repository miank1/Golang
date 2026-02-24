package main

import "fmt"

type User struct {
	Name string
}

func (u User) ChangeName(newName string) {
	u.Name = newName
}

func main() {

	u := User{Name: "Alex"}
	u.ChangeName("Farah")

	fmt.Println(u.Name)

}
