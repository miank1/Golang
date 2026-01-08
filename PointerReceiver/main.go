package main

import "fmt"

type User struct {
	Name string
}

func (u *User) SetName(name string) {
	u.Name = name
}

func main() {
	u := User{
		Name: "Ankit",
	}
	u.SetName("Rahul")
	fmt.Println(u)
}
