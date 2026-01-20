package main

import "fmt"

type MyError struct{}

func (e *MyError) Error() string {
	return "something went wrong"
}
func mightFail() error {
	var err *MyError = nil
	return err
}
func main() {
	err := mightFail()
	if err != nil {
		fmt.Println("Error occurred:")
	} else {
		fmt.Println("All good")
	}
}
