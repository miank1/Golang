package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func UpdateUser(u *User) {

	u.Name = u.Name + " Jr"
	u.Age = u.Age + 1
}

func main() {

	u := User{
		Name: "Alice",
		Age:  30,
	}

	UpdateUser(&u)

	fmt.Println(u)
}
