package main

import "fmt"

type User struct {
	name string
	age  int
}

func updateAge(u *User, newAge int) {
	u.age = newAge
}

func main() {
	user := User{
		name: "John",
		age:  26,
	}
	fmt.Println("Before:", user)

	updateAge(&user, 30)

	fmt.Println("After:", user)
}
