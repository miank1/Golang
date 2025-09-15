// Pipeline Pattern
package main

import (
	"fmt"
	"strings"
)

func UpperCase(str []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, s := range str {
			out <- strings.ToUpper(s)
		}
	}()
	return out
}

func ReverseString(in <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for s := range in {
			for i := len(s) - 1; i >= 0; i-- {
				ch <- string(s[i])
			}
		}

	}()
	return ch
}

func Print(ch <-chan string) {
	for msg := range ch {
		fmt.Printf("%s", msg)
	}
}

func main() {
	result := []string{"go", "is", "fun", "to", "learn"}
	UpperCase(result)
	fmt.Println(result)
	reversed := ReverseString(UpperCase(result))
	Print(reversed)

}
