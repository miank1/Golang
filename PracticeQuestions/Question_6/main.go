package main

import "fmt"

type Logger interface {
	Log(message string)
}

type FileLogger struct {
}

type ConsoleLogger struct {
}

func (fl FileLogger) Log(msg string) {
	fmt.Println(" :", msg)
}

func (cl ConsoleLogger) Log(msg string) {
	fmt.Println(" :", msg)
}

func ProcessLogger(l Logger, message string) {

}

func main() {

}
