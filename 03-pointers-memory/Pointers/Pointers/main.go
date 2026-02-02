package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func Update(val *int) {
	*val = *val + 10
}

func (u *User) HaveBirthday() {
	u.Age++
}

// Value Receiver copies the struct - Pointer receiver modify the original.

func (u User) ChangeName(name string) {
	u.Name = name // won't affect original
}

func (u *User) RealChangeName(name string) {
	u.Name = name // will change original
}

type Task struct {
	Title string
	Done  bool
}

func main() {
	var x int = 10
	var p *int = &x

	fmt.Println("Value of x is ", x)
	fmt.Println("Value of *p is ", *p)
	*p = 20
	fmt.Println("Value of *p after update is ", *p)

	x1 := 5
	Update(&x1)
	fmt.Println("Value of x1 is ", x1)

	u := User{Name: "Alice", Age: 25}
	p1 := &u

	fmt.Println(p1.Name)
	fmt.Println((*p1).Age)

	u = User{
		Name: "Bob",
		Age:  30,
	}

	u.ChangeName("John")
	u.RealChangeName("Alex") // This will change the name.

	u.HaveBirthday()
	fmt.Println(u.Name)
	fmt.Println(u.Age)

	tasks := []*Task{
		{Title: "Learn Go", Done: false},
		{Title: "Build Project", Done: false},
	}

	// Mark all the tasks as done
	for _, t := range tasks {
		t.Done = true
	}

	for _, t := range tasks {
		fmt.Println(t.Title, "Done?", t.Done)
	}

}
