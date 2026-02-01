package main

import "fmt"

type MyError struct {
	Message string
}

func (e MyError) Error() string {
	return e.Message
}

func doSomething() error {
	return MyError{Message: "something went wrong"}
}

type NotFoundError struct {
	Resource string
	ID       int
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s with id %d not found", e.Resource, e.ID)
}
