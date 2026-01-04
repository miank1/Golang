// Interface embedding helps build complex behaviors by composing smaller, focused interfaces- without duplicating method definitions.

package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(string)
}

// This is composition

type ReaderWriter interface {
	Reader
	Writer
}

type File struct{}

func (f File) Read() string {
	return "reading file"
}

func (f File) Write(data string) {
	fmt.Println("writing:", data)
}

func process(rw ReaderWriter) {
	fmt.Println(rw.Read())
	rw.Write("hello")
}

func main() {

	f := File{}
	process(f)
}
